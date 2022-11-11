# netbackup-exporter

Netbackup admin job metrics exporter to prometheus
![Screenshot 2022-11-11 at 09 08 31](https://user-images.githubusercontent.com/9497591/201261575-a1f05ef9-6ccd-45ac-8a4f-82bc351e2a90.png)

## Getting started

https://sort.veritas.com/public/documents/nbu/8.1.2/windowsandunix/productguides/html/index/#getting-started

1. Download docker-compose.yaml:
```bash
curl -fsSL https://raw.githubusercontent.com/grafana/oncall/dev/docker-compose.yml -o docker-compose.yml
```

2. Set variables, use [NetBackup API Getting Started](https://sort.veritas.com/public/documents/nbu/8.1.2/windowsandunix/productguides/html/index/#getting-started):
```bash
echo "NBU.APIKEY_TEST=MY_API_KEY_OR_JWT_TOKEN
NBU.MASTERSERVER=https://mymasterserver.com/netbackup
NBU.CACERT=----BEGIN CERTIFICATE-----mymasterserver_cacert----END CERTIFICATE-----" > .env_exporter
```

3. Launch exporter:
```bash
docker-compose --env-file .env_exporter -f docker-compose.yml up -d
```

## Overview
```bash
Usage of ./netbackup-exporter:
      --nbu.CACert string              CA certificate from the master server using the GET /security/cacert API
      --nbu.apiKey string              API key for NBU the /webui/security/api-keys
      --nbu.http.insecureSkipVerify    controls whether a client verifies the server's certificate chain and host name
      --nbu.http.reqTimeout duration   netBackup api request http timeout (default 11s)
      --nbu.jobsGetFilter string       Gets the list of jobs based on specified filters
      --nbu.jobsLast duration          retrieve the information on the backup jobs behind aggregation time (default 12h0m0s)
      --nbu.jobsPageLimit int          The number of records on one page after the offset (default 10)
      --nbu.masterServer string        netBackup master server base url
      --port string                    listen metrics port (default "9100")
pflag: help requested
```

## Metrics

Metric Name                            | Type      | Description
----------------------------------------|-----------|-------------------------------------------
 nbu_jobs_last_hours_ago_total          | Gauge     | Total number of netbackup jobs in a few hours
 nbu_jobs_kilobytes_transferred_total   | Counter   | The total kilobytes transferred of netbackup job
 nbu_jobs_elapsed_time_bucket           | Histogram | The elapsed time of netbackup jobs
