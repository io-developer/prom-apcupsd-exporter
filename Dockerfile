FROM ubuntu:22.04

ENV EXPORTER_LOG_LEVEL "info"
ENV EXPORTER_HOST "0.0.0.0"
ENV EXPORTER_PORT "8001"
ENV EXPORTER_APCUPSD_HOST "127.0.0.1"
ENV EXPORTER_APCUPSD_PORT "3551"
ENV EXPORTER_APCACCESS_PATH "/sbin/apcaccess"
ENV EXPORTER_APCACCESS_FLOOD_LIMIT_SECONDS "0.5"
ENV EXPORTER_APCACCESS_ERROR_IGNORING_SECONDS "120"
ENV EXPORTER_APCUPSD_STARTUP_WAIT_SECONDS "60"
ENV EXPORTER_PERSISTENT_COLLECT_INTERVAL_SECONDS "30"
ENV EXPORTER_DEFAULT_STATE_JSON "'{}'"

ENV \
  LANGUAGE=en_US.UTF-8 \
  LANG=en_US.utf8 \
  LC_ALL=en_US.UTF-8 \
  LC_CTYPE=en_US.UTF-8 \
  TERM=xterm

RUN \
  export DEBIAN_FRONTEND=noninteractive \
  && apt-get update \
  && apt-get install -y --no-install-recommends apcupsd \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ADD bin/prom-apcupsd-exporter /prom-apcupsd-exporter

ADD docker_cmd.sh /docker_cmd.sh
RUN chmod +x /docker_cmd.sh

EXPOSE ${EXPORTER_PORT}

CMD ["/docker_cmd.sh"]
