package server

import (
	"fmt"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/model"
	"math"
	"net/http"
	"time"
)

// metricsInit ..
func eventsInit() {
	http.HandleFunc("/event/commfailure", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("commfailure", w, r)
	})
	http.HandleFunc("/event/commok", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("commok", w, r)
	})

	http.HandleFunc("/event/startselftest", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("startselftest", w, r)
	})
	http.HandleFunc("/event/endselftest", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("endselftest", w, r)
	})

	http.HandleFunc("/event/powerout", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("powerout", w, r)
	})
	http.HandleFunc("/event/mainsback", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("mainsback", w, r)
	})

	http.HandleFunc("/event/onbattery", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("onbattery", w, r)
	})
	http.HandleFunc("/event/offbattery", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("offbattery", w, r)
	})
	http.HandleFunc("/event/battdetach", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("battdetach", w, r)
	})
	http.HandleFunc("/event/battattach", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("battattach", w, r)
	})
	http.HandleFunc("/event/changeme", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("changeme", w, r)
	})

	http.HandleFunc("/event/failing", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("failing", w, r)
	})
	http.HandleFunc("/event/timeout", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("timeout", w, r)
	})
	http.HandleFunc("/event/loadlimit", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("loadlimit", w, r)
	})
	http.HandleFunc("/event/runlimit", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("runlimit", w, r)
	})
	http.HandleFunc("/event/doshutdown", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("doshutdown", w, r)
	})
	http.HandleFunc("/event/annoyme", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("annoyme", w, r)
	})
	http.HandleFunc("/event/emergency", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("emergency", w, r)
	})
	http.HandleFunc("/event/remotedown", func(w http.ResponseWriter, r *http.Request) {
		eventsHandle("remotedown", w, r)
	})

}

func eventsHandle(event string, w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandle")
	WsBroadcastData(map[string]interface{}{
		"type":       "event",
		"event_type": event,
	})
	w.Write([]byte("ok"))

	ev := model.Event{
		Ts:   time.Now(),
		Name: event,
	}

	m := collector.GetModel()
	if event == "offbattery" {
		ev = handleEventOffbattery(m, ev)
	}

	m.AddEvent(ev)
	collector.Collect(metric.CollectOpts{
		PreventFlood: false,
	})
}

func handleEventOffbattery(m *model.Model, ev model.Event) model.Event {
	ondate := m.State.UpsTransferOnBatteryDate
	offdate := m.State.UpsTransferOffBatteryDate

	deltaOn := offdate.Sub(ondate).Seconds()
	deltaNow := offdate.Sub(time.Now()).Seconds()

	if deltaOn < 0 || math.Abs(deltaNow) < 30 {
		ev.Text = fmt.Sprintf(
			"{duration: \"%s\"}",
			fmtDuration(m.State.GetLastUpsOnBatteryDuration()),
		)
	}
	return ev
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
