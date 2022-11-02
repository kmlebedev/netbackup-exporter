FROM registry.tochka-tech.com/proxy_docker-io/library/golang:1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=http://nexus.bank24.int/repository/golang-public/ \
    GOSUMDB='sum.golang.org http://nexus.bank24.int/repository/golang-proxy-sum-golang-org/'

RUN mkdir -p /go/src/gitlab.tochka-tech.com/devexp/oci/netbackup-exporter
WORKDIR /go/src/gitlab.tochka-tech.com/devexp/oci/netbackup-exporter

COPY *.go ./go.mod ./go.sum ./
COPY ./nbu-admin-api ./nbu-admin-api
RUN CGO_ENABLED=0 GOOS=linux go install

FROM registry.tochka-tech.com/oci_rhel/ubi-minimal:v8

USER nobody
WORKDIR /app

COPY --from=builder /go/bin/netbackup-exporter ./netbackup_exporter
ENTRYPOINT ["/app/netbackup_exporter"]
