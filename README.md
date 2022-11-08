# netbackup-exporter

Netbackup admin job metrics exporter to prometheus

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