package metric

import "github.com/prometheus/common/promlog"

// Logger - final prometheus logger
var Logger = promlog.New(&promlog.Config{})
