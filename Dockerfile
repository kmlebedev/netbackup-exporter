FROM golang:1.19-alpine as builder

RUN mkdir -p /go/src/github.com/kmlebedev/netbackup-exporter
WORKDIR /go/src/github.com/kmlebedev/netbackup-exporter

COPY *.go ./go.mod ./go.sum ./
COPY ./nbu-admin-api ./nbu-admin-api
RUN CGO_ENABLED=0 GOOS=linux go install

FROM alpine

USER nobody
WORKDIR /app

COPY --from=builder /go/bin/netbackup-exporter ./netbackup_exporter
ENTRYPOINT ["/app/netbackup_exporter"]
