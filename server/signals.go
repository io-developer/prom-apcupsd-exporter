package server

import (
	"net/http"
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/metric"
	"github.com/io-developer/prom-apcupsd-exporter/model"

	"github.com/go-kit/kit/log/level"
)

// metricsInit ..
func signalsInit() {
	signals := []model.Signal{
		model.SignalCommfailure,
		model.SignalCommok,
		model.SignalStartselftest,
		model.SignalEndselftest,
		model.SignalPowerout,
		model.SignalMainsback,
		model.SignalOnbattery,
		model.SignalOffbattery,
		model.SignalBattattach,
		model.SignalChangeme,
		model.SignalFailing,
		model.SignalTimeout,
		model.SignalLoadlimit,
		model.SignalDoshutdown,
		model.SignalAnnoyme,
		model.SignalRemotedown,
	}
	for _, signal := range signals {
		signalsRegisterEndpoint(signal)
	}
}

func signalsRegisterEndpoint(signal model.Signal) {
	http.HandleFunc("/signal/"+string(signal), func(w http.ResponseWriter, r *http.Request) {
		signalsHandle(signal, w, r)
	})
}

func signalsHandle(signal model.Signal, w http.ResponseWriter, r *http.Request) {
	level.Info(logger).Log("msg", "handle signal", "signal", signal)

	WsBroadcastData(map[string]interface{}{
		"type":   "signal",
		"signal": string(signal),
	})
	w.Write([]byte("ok"))

	collector.GetModel().AddEvent(model.Event{
		Ts:   time.Now(),
		Type: model.EventTypeSignal,
		Data: map[string]interface{}{
			"signal": signal,
		},
	})

	collector.Collect(metric.CollectOpts{
		PreventFlood: false,
		Signal:       signal,
	})
}
