package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/antihax/optional"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	swagger "gitlab.tochka-tech.com/devexp/oci/netbackup-exporter/nbu-admin-api"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NbuExporter struct {
	nbuAdminApiClient *swagger.APIClient
	jobMetricsDataInc map[string]map[string]int
}

var (
	nbuExporter      *NbuExporter
	nbuReqTimeout    time.Duration
	nbuJobsGetFilter string
	nbuJobsPageLimit optional.Int32
	jobsCounterVec   = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "nbu_jobs_total",
		Help: "The total number of netbackup jobs",
	}, []string{"state", "type", "policyType", "clientName", "status"})
)

func init() {
	flag.String("port", "", "listen metrics port")
	flag.String("nbu.masterServer", "", "netBackup master server base url")
	flag.String("nbu.apiKey", "", "The JSON Web Token (JWT) or API key for NBU")
	flag.Duration("nbu.reqTimeout", 30000, "netBackup api request timeout")
	flag.String("nbu.jobsGetFilter", "", "Gets the list of jobs based on specified filters")
	flag.Int("nbu.jobsPageLimit", 10, "The number of records on one page after the offset")
	flag.Parse()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		glog.Exit(err)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	nbuExporter = &NbuExporter{
		nbuAdminApiClient: swagger.NewAPIClient(&swagger.Configuration{
			BasePath:      viper.GetString("nbu.masterServer"),
			DefaultHeader: map[string]string{"Authorization": viper.GetString("nbu.apiKey")},
			UserAgent:     "gitlab.tochka-tech.com/devexp/oci/netbackup-exporter/1.0.0/go",
		}),
		jobMetricsDataInc: make(map[string]map[string]int),
	}
	nbuReqTimeout = viper.GetDuration("nbu.reqTimeout")
	nbuJobsGetFilter = viper.GetString("nbu.jobsGetFilter")
	nbuJobsPageLimit = optional.NewInt32(viper.GetInt32("nbu.jobsPageLimit"))
}

// implements prometheus.Collector.
func (e *NbuExporter) Describe(ch chan<- *prometheus.Desc) {
	// empty
}

// Todo one metric collect
func (e *NbuExporter) Collect(ch chan<- prometheus.Metric) {
	// Todo pass last colect time
	nbuJobsGetFilterEndTime := fmt.Sprintf("endTime gt %s", time.Now())
	var nbuJobsGetFilterOpt optional.String
	if nbuJobsGetFilter == "" {
		nbuJobsGetFilterOpt = optional.NewString(nbuJobsGetFilterEndTime)
	} else {
		nbuJobsGetFilterOpt = optional.NewString(fmt.Sprintf("%s and %s", nbuJobsGetFilter, nbuJobsGetFilterEndTime))
	}
	jobsPageOffset := optional.EmptyInt32()
	for {
		ctx, cancel := context.WithTimeout(context.Background(), nbuReqTimeout)
		jobs, _, err := e.nbuAdminApiClient.JobsApi.AdminJobsGet(ctx, &swagger.JobsApiAdminJobsGetOpts{
			Filter:     nbuJobsGetFilterOpt,
			PageLimit:  nbuJobsPageLimit,
			PageOffset: jobsPageOffset,
		})
		defer cancel()
		if err != nil {
			glog.Errorf("NbuExporter.AdminJobsGet: %v", err)
			break
		}
		for _, jobData := range jobs.Data {
			jobsCounterMetric := jobsCounterVec.WithLabelValues(
				jobData.Attributes.State,
				jobData.Attributes.JobType,
				jobData.Attributes.PolicyType,
				jobData.Attributes.ClientName,
				strconv.FormatInt(int64(jobData.Attributes.Status), 10),
			)
			jobsCounterMetric.Inc()
			ch <- jobsCounterMetric
		}
		jobsPageOffset = optional.NewInt32(jobs.Meta.Pagination.Page)
		if jobs.Meta.Pagination.Page == jobs.Meta.Pagination.Last {
			break
		}
	}
}

func main() {
	registry := prometheus.NewRegistry()
	registry.MustRegister(nbuExporter)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	addr := fmt.Sprintf(":%s", viper.GetString("port"))
	glog.Infof("Start listen %s", addr)
	http.Handle("/metrics", handler)
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok, %q", html.EscapeString(r.URL.Path))
	})
	if err := http.ListenAndServe(addr, nil); err != nil {
		glog.Fatalf("http.ListenAndServer: %v", err)
	}
}
