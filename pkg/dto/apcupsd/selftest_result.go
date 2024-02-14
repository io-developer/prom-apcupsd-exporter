package apcupsd

type SelftestResult uint8

const (
	SELFTEST_RESULT__NA = SelftestResult(iota)
	SELFTEST_RESULT__NONE
	SELFTEST_RESULT__FAILED
	SELFTEST_RESULT__WARNING
	SELFTEST_RESULT__IN_PROGRESS
	SELFTEST_RESULT__PASSED
	SELFTEST_RESULT__FAILCAP
	SELFTEST_RESULT__FAILLOAD
	SELFTEST_RESULT__UNKNOWN
)

var SELFTEST_RESULT_TO_NAME = map[SelftestResult]string{
	SELFTEST_RESULT__NA:          "Not supported",
	SELFTEST_RESULT__NONE:        "No test results available",
	SELFTEST_RESULT__FAILED:      "Test failed",
	SELFTEST_RESULT__WARNING:     "Warning",
	SELFTEST_RESULT__IN_PROGRESS: "In progress",
	SELFTEST_RESULT__PASSED:      "Battery OK",
	SELFTEST_RESULT__FAILCAP:     "Test failed -- insufficient battery capacity",
	SELFTEST_RESULT__FAILLOAD:    "Test failed -- battery overloaded",
	SELFTEST_RESULT__UNKNOWN:     "Unknown",
}

var SELFTEST_RESULT_TO_SHORT_NAME = map[SelftestResult]string{
	SELFTEST_RESULT__NA:          "NA",
	SELFTEST_RESULT__NONE:        "none",
	SELFTEST_RESULT__FAILED:      "failed",
	SELFTEST_RESULT__WARNING:     "warning",
	SELFTEST_RESULT__IN_PROGRESS: "in_progress",
	SELFTEST_RESULT__PASSED:      "passed",
	SELFTEST_RESULT__FAILCAP:     "failcap",
	SELFTEST_RESULT__FAILLOAD:    "failload",
	SELFTEST_RESULT__UNKNOWN:     "unknown",
}

var SELFTEST_RESULT_FROM_NAME = map[string]SelftestResult{
	"Not supported":             SELFTEST_RESULT__NA,
	"No test results available": SELFTEST_RESULT__NONE,
	"Test failed":               SELFTEST_RESULT__FAILED,
	"Warning":                   SELFTEST_RESULT__WARNING,
	"In progress":               SELFTEST_RESULT__IN_PROGRESS,
	"Battery OK":                SELFTEST_RESULT__PASSED,
	"Test failed -- insufficient battery capacity": SELFTEST_RESULT__FAILCAP,
	"Test failed -- battery overloaded":            SELFTEST_RESULT__FAILLOAD,
	"Unknown":                                      SELFTEST_RESULT__UNKNOWN,
}

var SELFTEST_RESULT_FROM_SHORT_NAME = map[string]SelftestResult{
	"NA":          SELFTEST_RESULT__NA,
	"none":        SELFTEST_RESULT__NONE,
	"failed":      SELFTEST_RESULT__FAILED,
	"warning":     SELFTEST_RESULT__WARNING,
	"in_progress": SELFTEST_RESULT__IN_PROGRESS,
	"passed":      SELFTEST_RESULT__PASSED,
	"failcap":     SELFTEST_RESULT__FAILCAP,
	"failload":    SELFTEST_RESULT__FAILLOAD,
	"unknown":     SELFTEST_RESULT__UNKNOWN,
}
