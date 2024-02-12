package apcupsd

type SelftestResult uint8

const (
	SELFTEST_RESULT__NA         = SelftestResult(0)
	SELFTEST_RESULT__NONE       = SelftestResult(1)
	SELFTEST_RESULT__PASSED     = SelftestResult(2)
	SELFTEST_RESULT__FAILCAP    = SelftestResult(3)
	SELFTEST_RESULT__FAILED     = SelftestResult(4)
	SELFTEST_RESULT__INPROGRESS = SelftestResult(5)
	SELFTEST_RESULT__WARNING    = SelftestResult(6)
	SELFTEST_RESULT__UNKNOWN    = SelftestResult(7)
)
