FROM ubuntu:18.04

ENV \
  LANGUAGE=en_US.UTF-8 \
  LANG=en_US.utf8 \
  LC_ALL=en_US.UTF-8 \
  LC_CTYPE=en_US.UTF-8 \
  TERM=xterm

RUN \
  export DEBIAN_FRONTEND=noninteractive \
  && apt-get update \
  && apt-get install -y --no-install-recommends net-tools iputils-ping curl apcupsd \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ADD bin/prom-apcupsd-exporter /prom-apcupsd-exporter
RUN chmod +x /prom-apcupsd-exporter

EXPOSE 8001

ENTRYPOINT ["/prom-apcupsd-exporter"]
