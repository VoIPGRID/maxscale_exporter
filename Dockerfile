FROM        quay.io/prometheus/busybox:latest
MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>

COPY maxscale_exporter /bin/maxscale_exporter

ENTRYPOINT  ["/bin/maxscale_exporter"]
USER        nobody
EXPOSE      9195
