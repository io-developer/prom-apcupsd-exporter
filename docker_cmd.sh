#!/bin/sh

/prom-apcupsd-exporter                                              \
  -logLevel="$EXPORTER_LOG_LEVEL"                                   \
  -listen="$EXPORTER_HOST:$EXPORTER_PORT"                           \
  -apcupsd="$EXPORTER_APCUPSD_HOST:$EXPORTER_APCUPSD_PORT"          \
  -apcaccess="$EXPORTER_APCACCESS_PATH"                             \
  -floodLimit="$EXPORTER_APCACCESS_FLOOD_LIMIT_SECONDS"             \
  -errorIgnoreTime="$EXPORTER_APCACCESS_ERROR_IGNORING_SECONDS"     \
  -apcupsdStartSkip="$EXPORTER_APCUPSD_STARTUP_WAIT_SECONDS"        \
  -collectInterval="$EXPORTER_PERSISTENT_COLLECT_INTERVAL_SECONDS"  \
  -defaultModelState="$EXPORTER_DEFAULT_STATE_JSON"