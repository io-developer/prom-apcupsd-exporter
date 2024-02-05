FROM ubuntu:20.04

ENV EXPORTER_HOST="0.0.0.0"
ENV EXPORTER_PORT="8001"
ENV EXPORTER_APCUPSD_HOST="127.0.0.1"
ENV EXPORTER_APCUPSD_PORT="3551"
ENV EXPORTER_APCACCESS_PATH="/sbin/apcaccess"
ENV EXPORTER_APCACCESS_FLOOD_LIMIT_SECONDS="0.5"
ENV EXPORTER_APCACCESS_ERROR_IGNORING_SECONDS="120"
ENV EXPORTER_APCUPSD_STARTUP_WAIT_SECONDS="60"
ENV EXPORTER_PERSISTENT_COLLECT_INTERVAL_SECONDS="30"
ENV EXPORTER_DEFAULT_STATE_JSON=""

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

ADD bin/prom-smartctl-exporter /prom-smartctl-exporter

EXPOSE ${EXPORTER_PORT}

CMD [ \
  "/prom-apcupsd-exporter", \
  "-listen", ${EXPORTER_HOST}:${EXPORTER_PORT}, \
  "-apcupsd", ${EXPORTER_APCUPSD_HOST}:${EXPORTER_APCUPSD_PORT} \
  "-apcaccess", ${EXPORTER_APCACCESS_PATH} \
  "-floodLimit", ${EXPORTER_APCACCESS_FLOOD_LIMIT_SECONDS} \
  "-errorIgnoreTime", ${EXPORTER_APCACCESS_ERROR_IGNORING_SECONDS} \
  "-apcupsdStartSkip", ${EXPORTER_APCUPSD_STARTUP_WAIT_SECONDS} \
  "-collectInterval", ${EXPORTER_PERSISTENT_COLLECT_INTERVAL_SECONDS} \
  "-defStateJSON", ${EXPORTER_DEFAULT_STATE_JSON} \
]
