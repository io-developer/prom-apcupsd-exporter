package main

import (
	"flag"
	"local/apcupsd_exporter/server"
)

func main() {
	listenArg := flag.String("listen", "0.0.0.0:8001", "ip:port")
	upsArg := flag.String("ups", "127.0.0.1:3551", "apcupsd host:port")
	apcaccessArg := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	flag.Parse()

	server.ListenAndServe(*listenArg, *upsArg, *apcaccessArg)
}
