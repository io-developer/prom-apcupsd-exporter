package apcupsd

type Sensivity uint8

const (
	SENSIVITY__NA = Sensivity(iota)
	SENSIVITY__AUTO
	SENSIVITY__LOW
	SENSIVITY__MEDIUM
	SENSIVITY__HIGH
	SENSIVITY__UNKNOWN
)

var SENSIVITY_TO_NAME = map[Sensivity]string{
	SENSIVITY__NA:      "NA",
	SENSIVITY__AUTO:    "Auto Adjust",
	SENSIVITY__LOW:     "Low",
	SENSIVITY__MEDIUM:  "Medium",
	SENSIVITY__HIGH:    "High",
	SENSIVITY__UNKNOWN: "Unknown",
}

var SENSIVITY_TO_SHORT_NAME = map[Sensivity]string{
	SENSIVITY__NA:      "NA",
	SENSIVITY__AUTO:    "auto",
	SENSIVITY__LOW:     "low",
	SENSIVITY__MEDIUM:  "medium",
	SENSIVITY__HIGH:    "high",
	SENSIVITY__UNKNOWN: "unknown",
}

var SENSIVITY_FROM_NAME = map[string]Sensivity{
	"NA":          SENSIVITY__NA,
	"Auto Adjust": SENSIVITY__AUTO,
	"Low":         SENSIVITY__LOW,
	"Medium":      SENSIVITY__MEDIUM,
	"High":        SENSIVITY__HIGH,
	"Unknown":     SENSIVITY__UNKNOWN,
}

var SENSIVITY_FROM_SHORT_NAME = map[string]Sensivity{
	"NA":      SENSIVITY__NA,
	"auto":    SENSIVITY__AUTO,
	"low":     SENSIVITY__LOW,
	"medium":  SENSIVITY__MEDIUM,
	"high":    SENSIVITY__HIGH,
	"unknown": SENSIVITY__UNKNOWN,
}
