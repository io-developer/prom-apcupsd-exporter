package server

import "net/http"

// metricsInit ..
func eventsInit() {
	http.HandleFunc("/event/commfailure", eventsHandleCommfailure)
	http.HandleFunc("/event/commok", eventsHandleCommok)

	http.HandleFunc("/event/startselftest", eventsHandleStartselftest)
	http.HandleFunc("/event/endselftest", eventsHandleEndselftest)

	http.HandleFunc("/event/powerout", eventsHandlePowerout)
	http.HandleFunc("/event/mainsback", eventsHandleMainsback)

	http.HandleFunc("/event/onbattery", eventsHandleOnbattery)
	http.HandleFunc("/event/offbattery", eventsHandleOffbattery)
	http.HandleFunc("/event/battdetach", eventsHandleBattdetach)
	http.HandleFunc("/event/battattach", eventsHandleBattattach)
	http.HandleFunc("/event/changeme", eventsHandleChangeme)

	http.HandleFunc("/event/failing", eventsHandleFailing)
	http.HandleFunc("/event/timeout", eventsHandleTimeout)
	http.HandleFunc("/event/loadlimit", eventsHandleLoadlimit)
	http.HandleFunc("/event/runlimit", eventsHandleRunlimit)
	http.HandleFunc("/event/doshutdown", eventsHandleDoshutdown)
	http.HandleFunc("/event/annoyme", eventsHandleAnnoyme)
	http.HandleFunc("/event/emergency", eventsHandleEmergency)
	http.HandleFunc("/event/remotedown", eventsHandleRemotedown)

}

func eventsHandleCommfailure(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleCommfailure")
}

func eventsHandleCommok(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleCommok")
}

func eventsHandleStartselftest(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleStartselftest")
}

func eventsHandleEndselftest(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleEndselftest")
}

func eventsHandlePowerout(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandlePowerout")
}

func eventsHandleMainsback(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleMainsback")
}

func eventsHandleOnbattery(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleOnbattery")
}

func eventsHandleOffbattery(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleOffbattery")
}

func eventsHandleBattdetach(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleBattdetach")
}

func eventsHandleBattattach(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleBattattach")
}

func eventsHandleChangeme(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleChangeme")
}

func eventsHandleFailing(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleFailing")
}

func eventsHandleTimeout(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleTimeout")
}

func eventsHandleLoadlimit(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleLoadlimit")
}

func eventsHandleRunlimit(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleRunlimit")
}

func eventsHandleDoshutdown(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleDoshutdown")
}

func eventsHandleAnnoyme(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleAnnoyme")
}

func eventsHandleEmergency(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleEmergency")
}

func eventsHandleRemotedown(w http.ResponseWriter, r *http.Request) {
	logger.Log("msg", "eventsHandleRemotedown")
}
