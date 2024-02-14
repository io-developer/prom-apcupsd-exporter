package apcupsd

type Mode uint8

const (
	MODE__NA = Mode(iota)
	MODE__STAND_ALONE
	MODE__SHARE_SLAVE
	MODE__SHARE_MASTER
)

var MODE_TO_NAME = map[Mode]string{
	MODE__NA:           "*invalid-ups-class*",
	MODE__STAND_ALONE:  "Stand Alone",
	MODE__SHARE_SLAVE:  "ShareUPS Slave",
	MODE__SHARE_MASTER: "ShareUPS Master",
}

var MODE_TO_SHORT_NAME = map[Mode]string{
	MODE__NA:           "NA",
	MODE__STAND_ALONE:  "standalone",
	MODE__SHARE_SLAVE:  "shareslave",
	MODE__SHARE_MASTER: "sharemaster",
}

var MODE_FROM_NAME = map[string]Mode{
	"*invalid-ups-class*": MODE__NA,
	"Stand Alone":         MODE__STAND_ALONE,
	"ShareUPS Slave":      MODE__SHARE_SLAVE,
	"ShareUPS Master":     MODE__SHARE_MASTER,
}

var MODE_FROM_SHORT_NAME = map[string]Mode{
	"NA":          MODE__NA,
	"standalone":  MODE__STAND_ALONE,
	"shareslave":  MODE__SHARE_SLAVE,
	"sharemaster": MODE__SHARE_MASTER,
}
