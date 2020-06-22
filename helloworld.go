package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
)

var (
	gaugeFunc = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "aaa_gauge_func",
	}, func() float64 {
		promLog.Infoln("GET METRIC")
		return float64(time.Now().Second())
	})

	gaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aaa_gauge_vec",
	}, []string{"now"})

	histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:        "aaa_histogram",
		Buckets:     prometheus.LinearBuckets(0, 6, 10),
		ConstLabels: prometheus.Labels{"foo": "bar", "hello": "kitty"},
	})

	summary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "aaa_summary",
		MaxAge:     10 * time.Second,
		AgeBuckets: 1,
	})

	summaryQuantile = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "aaa_summary_quantile_minute_div2",
		Objectives: map[float64]float64{
			1.0: 0,
		},
		MaxAge:     1 * time.Minute,
		AgeBuckets: 3,
	})

	summaryVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "aaa_summary_vec",
		Objectives: map[float64]float64{
			1.0: 0,
		},
		MaxAge:     5 * time.Second,
		AgeBuckets: 3,
	}, []string{"now"})
)

var lastUpdateTs = int64(0)

func updateMetrics() {
	nowTs := time.Now().UnixNano()
	promLog.Infoln("updateMetrics()", nowTs-lastUpdateTs)

	if nowTs-lastUpdateTs < 500_000_000 {
		promLog.Infoln("TOO FAST")
		return
	}
	lastUpdateTs = nowTs

	gaugeVec.With(prometheus.Labels{"now": "hour"}).Set(float64(time.Now().Hour()))
	gaugeVec.With(prometheus.Labels{"now": "minute"}).Set(float64(time.Now().Minute()))
	gaugeVec.With(prometheus.Labels{"now": "second"}).Set(float64(time.Now().Second()))

	secDivs := make([]float64, 60)
	for i := range secDivs {
		if i > 0 && time.Now().Second()%i == 0 {
			secDivs[i] = 1
		} else {
			secDivs[i] = 0
		}
	}

	minDivs := make([]float64, 60)
	for i := range minDivs {
		if i > 0 && time.Now().Minute()%i == 0 {
			minDivs[i] = 1
		} else {
			minDivs[i] = 0
		}
	}

	gaugeVec.With(prometheus.Labels{"now": "s3=0"}).Set(secDivs[3])
	gaugeVec.With(prometheus.Labels{"now": "s7=0"}).Set(secDivs[7])
	gaugeVec.With(prometheus.Labels{"now": "s17=0"}).Set(secDivs[17])
	gaugeVec.With(prometheus.Labels{"now": "s37=0"}).Set(secDivs[37])

	histogram.Observe(3 * secDivs[3])
	histogram.Observe(7 * secDivs[7])
	histogram.Observe(17 * secDivs[17])
	histogram.Observe(37 * secDivs[37])

	for _, i := range []int{3, 7, 17, 37} {
		if time.Now().Second()%i == 0 {
			summary.Observe(float64(i))
		}
	}

	summaryQuantile.Observe(minDivs[2])

	summaryVec.With(prometheus.Labels{"now": "s3=0"}).Observe(secDivs[3])
	summaryVec.With(prometheus.Labels{"now": "s7=0"}).Observe(secDivs[7])
	summaryVec.With(prometheus.Labels{"now": "s13=0"}).Observe(secDivs[17])
	summaryVec.With(prometheus.Labels{"now": "s37=0"}).Observe(secDivs[37])
}

func updateMetricsLoop() {
	for true {
		promLog.Infoln("updateMetricsLoop()")

		updateMetrics()
		time.Sleep(10 * time.Second)
	}
}

func main() {
	prometheus.MustRegister(gaugeFunc)
	prometheus.MustRegister(gaugeVec)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)
	prometheus.MustRegister(summaryQuantile)
	prometheus.MustRegister(summaryVec)

	go updateMetricsLoop()

	promHandler := promhttp.Handler()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promLog.Infoln("/metrics HANDLE")

		updateMetrics()
		promHandler.ServeHTTP(w, r)
	})

	//	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:7123", nil)
}
