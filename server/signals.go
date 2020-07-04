package server

import (
	"local/apcupsd_exporter/metric"
	"net/http"
)

// metricsInit ..
func signalsInit() {
	http.HandleFunc("/signal/commfailure", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("commfailure", w, r)
	})
	http.HandleFunc("/signal/commok", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("commok", w, r)
	})

	http.HandleFunc("/signal/startselftest", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("startselftest", w, r)
	})
	http.HandleFunc("/signal/endselftest", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("endselftest", w, r)
	})

	http.HandleFunc("/signal/powerout", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("powerout", w, r)
	})
	http.HandleFunc("/signal/mainsback", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("mainsback", w, r)
	})

	http.HandleFunc("/signal/onbattery", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("onbattery", w, r)
	})
	http.HandleFunc("/signal/offbattery", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("offbattery", w, r)
	})
	http.HandleFunc("/signal/battdetach", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("battdetach", w, r)
	})
	http.HandleFunc("/signal/battattach", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("battattach", w, r)
	})
	http.HandleFunc("/signal/changeme", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("changeme", w, r)
	})

	http.HandleFunc("/signal/failing", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("failing", w, r)
	})
	http.HandleFunc("/signal/timeout", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("timeout", w, r)
	})
	http.HandleFunc("/signal/loadlimit", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("loadlimit", w, r)
	})
	http.HandleFunc("/signal/runlimit", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("runlimit", w, r)
	})
	http.HandleFunc("/signal/doshutdown", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("doshutdown", w, r)
	})
	http.HandleFunc("/signal/annoyme", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("annoyme", w, r)
	})
	http.HandleFunc("/signal/emergency", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("emergency", w, r)
	})
	http.HandleFunc("/signal/remotedown", func(w http.ResponseWriter, r *http.Request) {
		signalsHandle("remotedown", w, r)
	})

}

func signalsHandle(signal string, w http.ResponseWriter, r *http.Request) {
	WsBroadcastData(map[string]interface{}{
		"type":   "signal",
		"signal": signal,
	})
	w.Write([]byte("ok"))

	collector.Collect(metric.CollectOpts{
		PreventFlood: false,
	})
}
