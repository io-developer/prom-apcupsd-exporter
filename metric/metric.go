package metric

import "github.com/prometheus/client_golang/prometheus"

// Metrics
var (
	BatteryCharge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "battery_charge",
		Help: "Current battery charge (percent)",
	})

	UpsStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ups_status",
		Help: "Current UPS Status (0=Calibration, 1=SmartTrim, 2=SmartBoost, 3=Online, 4=OnBattery, 5=Overloaded, 6=LowBattery, 7=ReplaceBattery, 8=OnBypass, 9=Off, 10=Charging, 11=Discharging)",
	})
)

// Register func
func Register() {
	prometheus.MustRegister(BatteryCharge)
	prometheus.MustRegister(UpsStatus)
}
