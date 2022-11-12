package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/antihax/optional"
	swagger "github.com/kmlebedev/netbackup-exporter/nbu-admin-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	glog "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"html"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	promNamespace = "nbu"
	promSubsystem = "jobs"
	timeZero      = "1970-01-01T00:00:00.000Z"
	stateActive   = "ACTIVE"
	stateDone     = "DONE"
)

type NbuExporter struct {
	nbuAdminApiClient *swagger.APIClient
	lastCollectTime   time.Time
	jobMetricsDataInc map[string]map[string]int
}

var (
	nbuExporter            *NbuExporter
	nbuHttpClinet          *http.Client
	nbuJobsGetFilter       string
	nbuJobsPageLimit       optional.Int32
	nbuJobsLastAggrTime    time.Duration
	nbuJobsLastReqInterval time.Duration
	nbuJobsLast            = make(map[Lables]uint64)
	nbuJobsLastCh          = make(chan bool)
	jobsKbytesTransferred  = make(map[int32]int32)
	jobLables              = []string{"state", "jobType", "policyType", "policyName", "clientName", "status"}
	jobsTotalLastHoursAgo  = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Subsystem: promSubsystem,
		Name:      "last_hours_ago_total",
		Help:      "Total number of netbackup jobs in a few hours",
	}, jobLables)
	jobsElapsedTimeHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: promNamespace,
		Subsystem: promSubsystem,
		Name:      "elapsed_time",
		Help:      "The elapsed time of netbackup jobs",
		Buckets:   []float64{1, 30, 60, 120, 300, 600, 720, 900, 1200, 1800, 2400, 3000, 3600, 5400, 7200, 10800, 140400},
	}, jobLables)
	jobsKilobytesTransferredVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: promNamespace,
		Subsystem: promSubsystem,
		Name:      "kilobytes_transferred_total",
		Help:      "The total kilobytes transferred of netbackup jobs",
	}, jobLables)
	jobsTransferRateVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Subsystem: promSubsystem,
		Name:      "kilobytes_transfer_rate",
		Help:      "The transfer rate of netbackup jobs",
	}, jobLables)
)

type Lables struct {
	State      string
	JobType    string
	PolicyType string
	PolicyName string
	ClientName string
	Status     string
}

func init() {
	flag.String("port", "9100", "listen metrics port")
	flag.Duration("nbu.jobsLast.AggrTime", 12*time.Hour, "retrieve the information on the backup jobs behind aggregation time")
	flag.Duration("nbu.jobsLast.ReqInterval", 10*time.Minute, "retrieve the information on the backup jobs for an interval")
	flag.String("nbu.masterServer", "", "netBackup master server base url")
	flag.String("nbu.apiKey", "", "API key for NBU the /webui/security/api-keys")
	flag.Duration("nbu.http.reqTimeout", 11*time.Second, "netBackup api request http timeout")
	flag.String("nbu.CACert", "", "CA certificate from the master server using the GET /security/cacert API")
	flag.Bool("nbu.http.insecureSkipVerify", false, "controls whether a client verifies the server's certificate chain and host name")
	flag.String("nbu.jobsGetFilter", "", "Gets the list of jobs based on specified filters")
	flag.Int("nbu.jobsPageLimit", 10, "The number of records on one page after the offset")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		glog.Fatal(err)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if configFile, ok := os.LookupEnv("CONFIG_FILE"); ok {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			glog.Fatalf("viper.ReadInConfig: %v", err)
		}
	}
	viper.AutomaticEnv()
	pflag.Parse()
	nbuJobsGetFilter = viper.GetString("nbu.jobsGetFilter")
	nbuJobsPageLimit = optional.NewInt32(viper.GetInt32("nbu.jobsPageLimit"))

	tlsConfig := tls.Config{}
	if nbuCaCert := viper.GetString("nbu.CACert"); nbuCaCert != "" {
		caCertPool := x509.NewCertPool()
		clientCACert, err := os.ReadFile(nbuCaCert)
		if err != nil {
			clientCACert = []byte(strings.ReplaceAll(nbuCaCert, "\\n", "\n"))
		}
		ok := caCertPool.AppendCertsFromPEM(clientCACert)
		glog.Debug("loaded %+v cert: %s", ok, string(clientCACert))
		tlsConfig.RootCAs = caCertPool
		tlsConfig.Certificates = []tls.Certificate{}
	}
	if viper.GetBool("nbu.http.insecureSkipVerify") {
		tlsConfig.InsecureSkipVerify = true
	}
	nbuHttpClinet = &http.Client{Transport: &http.Transport{TLSClientConfig: &tlsConfig}}
	nbuHttpClinet.Timeout = viper.GetDuration("nbu.reqTimeout")
	nbuJobsLastAggrTime = viper.GetDuration("nbu.jobsLast.AggrTime")
	nbuJobsLastReqInterval = viper.GetDuration("nbu.jobsLast.ReqInterval")

	nbuExporter = &NbuExporter{
		nbuAdminApiClient: swagger.NewAPIClient(&swagger.Configuration{
			BasePath:      viper.GetString("nbu.masterServer"),
			DefaultHeader: map[string]string{"Authorization": viper.GetString("nbu.apiKey")},
			UserAgent:     "gitlab.tochka-tech.com/devexp/oci/netbackup-exporter/1.0.0/go",
			HTTPClient:    nbuHttpClinet,
		}),
		jobMetricsDataInc: make(map[string]map[string]int),
		lastCollectTime:   time.Now(),
	}
}

func (e *NbuExporter) collectLastJobs(lastCollectTime time.Time) {
	if nbuJobsLastAggrTime == 0 || lastCollectTime.IsZero() {
		return
	}

	for lables, _ := range nbuJobsLast {
		nbuJobsLast[lables] = 0
	}
	e.GetAdminJobs(getJobsFilter(lastCollectTime.Add(-nbuJobsLastAggrTime)), func(jobData swagger.ManualBackupResponseData) {
		lables := Lables{
			State:      jobData.Attributes.State,
			JobType:    jobData.Attributes.JobType,
			PolicyType: jobData.Attributes.PolicyType,
			PolicyName: jobData.Attributes.PolicyName,
			ClientName: jobData.Attributes.ClientName,
			Status:     strconv.FormatInt(int64(jobData.Attributes.Status), 10),
		}
		nbuJobsLast[lables] += 1
	})
	jobsTotalLastHoursAgo.Reset()

	for lables, val := range nbuJobsLast {
		jobsTotalLastHoursAgo.WithLabelValues(
			lables.State,
			lables.JobType,
			lables.PolicyType,
			lables.PolicyName,
			lables.ClientName,
			lables.Status,
		).Set(float64(val))
	}
}

func (e *NbuExporter) startCollectLastJobs() {
	if nbuJobsLastReqInterval == 0 || nbuJobsLastAggrTime == 0 {
		return
	}

	ticker := time.NewTicker(nbuJobsLastReqInterval)
	lastCollectTime := time.Now()
	e.collectLastJobs(lastCollectTime)
	go func() {
		for {
			select {
			case <-nbuJobsLastCh:
				return
			case t := <-ticker.C:
				e.collectLastJobs(lastCollectTime)
				lastCollectTime = t
			}
		}
	}()
}

// implements prometheus.Collector.
func (e *NbuExporter) Describe(ch chan<- *prometheus.Desc) {
	jobsKilobytesTransferredVec.Describe(ch)
	jobsElapsedTimeHistogram.Describe(ch)
	jobsTransferRateVec.Describe(ch)
}

func (e *NbuExporter) GetAdminJobs(filter optional.String, fn func(jobData swagger.ManualBackupResponseData)) (jobCount uint32) {
	jobsPageOffset := optional.EmptyInt32()
	ctx := context.Background()
	for {
		jobs, resp, err := e.nbuAdminApiClient.JobsApi.AdminJobsGet(ctx, &swagger.JobsApiAdminJobsGetOpts{
			Filter:     filter,
			PageLimit:  nbuJobsPageLimit,
			PageOffset: jobsPageOffset,
		})
		if err != nil {
			glog.Errorf("NbuExporter.AdminJobsGet: %v resp: %+v, filter: %s", err, resp, filter.Value())
			break
		}
		for _, jobData := range jobs.Data {
			fn(jobData)
			jobCount++
		}
		if jobs.Meta == nil || jobs.Meta.Pagination == nil || jobs.Meta.Pagination.Last == 0 {
			break
		}
		if jobs.Meta.Pagination.Next == 0 {
			break
		}
		jobsPageOffset = optional.NewInt32(jobs.Meta.Pagination.Next)
	}
	return jobCount
}

func getJobsFilter(start time.Time) optional.String {
	filter := fmt.Sprintf("(endTime ge %s or endTime eq %s)", start.Format(time.RFC3339), timeZero)
	if nbuJobsGetFilter != "" {
		filter = fmt.Sprintf("%s and %s", filter, nbuJobsGetFilter)
	}
	return optional.NewString(filter)
}

func (e *NbuExporter) Collect(ch chan<- prometheus.Metric) {
	lastCollectTime := time.Now()

	jobCount := e.GetAdminJobs(getJobsFilter(e.lastCollectTime), func(jobData swagger.ManualBackupResponseData) {
		if jobData.Attributes.KilobytesTransferred > 0 {
			jobsKilobytesTransferredMetric := jobsKilobytesTransferredVec.WithLabelValues(
				jobData.Attributes.State,
				jobData.Attributes.JobType,
				jobData.Attributes.PolicyType,
				jobData.Attributes.PolicyName,
				jobData.Attributes.ClientName,
				strconv.FormatInt(int64(jobData.Attributes.Status), 10),
			)
			// active jobs
			if jobData.Attributes.State == stateDone || jobData.Attributes.EndTime.Unix() != 0 {
				if lastKbytesTransferred, ok := jobsKbytesTransferred[jobData.Attributes.JobId]; ok {
					if jobData.Attributes.KilobytesTransferred > lastKbytesTransferred {
						jobsKilobytesTransferredVec.WithLabelValues(
							stateActive,
							jobData.Attributes.JobType,
							jobData.Attributes.PolicyType,
							jobData.Attributes.PolicyName,
							jobData.Attributes.ClientName,
							strconv.FormatInt(int64(jobData.Attributes.Status), 10),
						).Add(float64(jobData.Attributes.KilobytesTransferred - lastKbytesTransferred))
					}
					delete(jobsKbytesTransferred, jobData.Attributes.JobId)
				}
				jobsKilobytesTransferredMetric.Add(float64(jobData.Attributes.KilobytesTransferred))
			} else {
				if lastKbytesTransferred, ok := jobsKbytesTransferred[jobData.Attributes.JobId]; ok {
					if jobData.Attributes.KilobytesTransferred > lastKbytesTransferred {
						jobsKilobytesTransferredMetric.Add(float64(jobData.Attributes.KilobytesTransferred - lastKbytesTransferred))
					}
				} else {
					jobsKilobytesTransferredMetric.Add(float64(jobData.Attributes.KilobytesTransferred))
				}
				jobsKbytesTransferred[jobData.Attributes.JobId] = jobData.Attributes.KilobytesTransferred
			}
		}

		if jobData.Attributes.TransferRate > 0 {
			jobsTransferRateVec.WithLabelValues(
				jobData.Attributes.State,
				jobData.Attributes.JobType,
				jobData.Attributes.PolicyType,
				jobData.Attributes.PolicyName,
				jobData.Attributes.ClientName,
				strconv.FormatInt(int64(jobData.Attributes.Status), 10),
			).Set(float64(jobData.Attributes.TransferRate))
		}

		if jobData.Attributes.EndTime.Unix() == 0 {
			return
		}

		jobsElapsedTimeMetric := jobsElapsedTimeHistogram.WithLabelValues(
			jobData.Attributes.State,
			jobData.Attributes.JobType,
			jobData.Attributes.PolicyType,
			jobData.Attributes.PolicyName,
			jobData.Attributes.ClientName,
			strconv.FormatInt(int64(jobData.Attributes.Status), 10),
		)
		if jobData.Attributes.ElapsedTime != "" {
			var h, m, s int
			if _, err := fmt.Sscanf(jobData.Attributes.ElapsedTime, "%2d:%2d:%2d", &h, &m, &s); err == nil {
				elapsedTime := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second
				jobsElapsedTimeMetric.Observe(elapsedTime.Seconds())
			} else {
				glog.Errorf("GetAdminJobs: invalid duration %+v: %+v", jobData.Attributes.ElapsedTime, err)
			}
		} else if jobData.Attributes.StartTime.Unix() != 0 && jobData.Attributes.EndTime.Unix() != 0 {
			if jobData.Attributes.EndTime.Unix() >= jobData.Attributes.StartTime.Unix() {
				jobsElapsedTimeMetric.Observe(jobData.Attributes.EndTime.Sub(jobData.Attributes.StartTime).Seconds())
			} else {
				glog.Errorf("GetAdminJobs: invalid StartTime > EndTime job: %+v", jobData.Attributes)
			}
		}
	})
	jobsKilobytesTransferredVec.Collect(ch)
	jobsTransferRateVec.Collect(ch)
	jobsElapsedTimeHistogram.Collect(ch)
	if nbuJobsLastAggrTime > 0 && nbuJobsLastReqInterval > 0 {
		jobsTotalLastHoursAgo.Collect(ch)
	}
	e.lastCollectTime = lastCollectTime
	if jobCount == 0 && len(jobsKbytesTransferred) > 0 {
		jobsKbytesTransferred = make(map[int32]int32)
	}
}

func main() {
	nbuExporter.startCollectLastJobs()
	defer close(nbuJobsLastCh)

	registry := prometheus.NewRegistry()
	registry.MustRegister(nbuExporter)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	addr := fmt.Sprintf(":%s", viper.GetString("port"))
	glog.Infof("Start listen %s", addr)
	http.Handle("/metrics", handler)
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "ok, %q", html.EscapeString(r.URL.Path)); err != nil {
			glog.Errorf("redy Fprintf: %v", err)
		}
	})
	if err := http.ListenAndServe(addr, nil); err != nil {
		glog.Fatalf("http.ListenAndServer: %v", err)
	}
}
