package apcupsd

type Sensivity uint8

const (
	SENSIVITY_NA      = Sensivity(0)
	SENSIVITY_LOW     = Sensivity(1)
	SENSIVITY_MEDIUM  = Sensivity(2)
	SENSIVITY_HIGH    = Sensivity(3)
	SENSIVITY_AUTO    = Sensivity(4)
	SENSIVITY_UNKNOWN = Sensivity(5)
)
