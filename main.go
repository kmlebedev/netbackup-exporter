package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/antihax/optional"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	glog "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.tochka-tech.com/devexp/oci/netbackup-exporter/nbu-admin-api"
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
)

type NbuExporter struct {
	nbuAdminApiClient *swagger.APIClient
	lastCollectTime   time.Time
	jobMetricsDataInc map[string]map[string]int
}

var (
	nbuExporter              *NbuExporter
	nbuHttpClinet            *http.Client
	nbuJobsGetFilter         string
	nbuJobsPageLimit         optional.Int32
	jobLables                = []string{"state", "type", "policyType", "clientName", "status"}
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
)

func init() {
	flag.String("port", "9100", "listen metrics port")
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

// implements prometheus.Collector.
func (e *NbuExporter) Describe(ch chan<- *prometheus.Desc) {
	jobsKilobytesTransferredVec.Describe(ch)
	jobsElapsedTimeHistogram.Describe(ch)
}

func (e *NbuExporter) Collect(ch chan<- prometheus.Metric) {
	// Todo pass last colect time
	endCollectTime := time.Now()
	nbuJobsGetFilterEndTime := fmt.Sprintf("startTime ge %s and endTime le %s",
		e.lastCollectTime.Format(time.RFC3339Nano), endCollectTime.Format(time.RFC3339Nano))
	e.lastCollectTime = endCollectTime
	var nbuJobsGetFilterOpt optional.String
	if nbuJobsGetFilter == "" {
		nbuJobsGetFilterOpt = optional.NewString(nbuJobsGetFilterEndTime)
	} else {
		nbuJobsGetFilterOpt = optional.NewString(fmt.Sprintf("%s and %s", nbuJobsGetFilter, nbuJobsGetFilterEndTime))
	}
	jobsPageOffset := optional.EmptyInt32()
	ctx := context.Background()
	for {
		jobs, resp, err := e.nbuAdminApiClient.JobsApi.AdminJobsGet(ctx, &swagger.JobsApiAdminJobsGetOpts{
			Filter:     nbuJobsGetFilterOpt,
			PageLimit:  nbuJobsPageLimit,
			PageOffset: jobsPageOffset,
		})
		if err != nil {
			glog.Errorf("NbuExporter.AdminJobsGet: %v resp: %+v, filter: %s", err, resp, nbuJobsGetFilterOpt.Value())
			break
		}
		for _, jobData := range jobs.Data {
			if jobData.Attributes.KilobytesTransferred > 0 {
				jobsKilobytesTransferredVec.WithLabelValues(
					jobData.Attributes.State,
					jobData.Attributes.JobType,
					jobData.Attributes.PolicyType,
					jobData.Attributes.ClientName,
					strconv.FormatInt(int64(jobData.Attributes.Status), 10),
				).Add(float64(jobData.Attributes.KilobytesTransferred))
			}

			if !jobData.Attributes.EndTime.IsZero() && !jobData.Attributes.StartTime.IsZero() {
				jobsElapsedTimeHistogram.WithLabelValues(
					jobData.Attributes.State,
					jobData.Attributes.JobType,
					jobData.Attributes.PolicyType,
					jobData.Attributes.ClientName,
					strconv.FormatInt(int64(jobData.Attributes.Status), 10),
				).Observe(jobData.Attributes.EndTime.Sub(jobData.Attributes.StartTime).Seconds())
			}
		}
		if jobs.Meta == nil || jobs.Meta.Pagination == nil || jobs.Meta.Pagination.Last == 0 {
			break
		}
		if jobs.Meta.Pagination.Next == 0 {
			break
		}
		jobsPageOffset = optional.NewInt32(jobs.Meta.Pagination.Next)
	}
	jobsKilobytesTransferredVec.Collect(ch)
	jobsElapsedTimeHistogram.Collect(ch)
}

func main() {
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
