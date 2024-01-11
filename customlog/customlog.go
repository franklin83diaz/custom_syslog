package customlog

import (
	"log"
	"log/syslog"
)

type LogType int

const (
	LOG_DEBUG LogType = iota
	LOG_INFO
	LOG_NOTICE
	LOG_WARNING
	LOG_ERR
	LOG_CRIT
	LOG_ALERT
	LOG_EMERG
)

// Logchpce is a function to write logs in syslog
//
// use: Logchpce("message", logchpce.LOG_INFO)
func Log(meesage string, t LogType) {

	nameprogram := "nameprogram"

	customlog, err := syslog.New(syslog.LOG_LOCAL6, nameprogram)
	if err != nil {
		log.Fatal(err)
	}
	switch t {
	case LOG_DEBUG:
		customlog.Debug(meesage)
	case LOG_INFO:
		customlog.Info(meesage)
	case LOG_NOTICE:
		customlog.Notice(meesage)
	case LOG_WARNING:
		customlog.Warning(meesage)
	case LOG_ERR:
		customlog.Err(meesage)
	case LOG_CRIT:
		customlog.Crit(meesage)
	case LOG_ALERT:
		customlog.Alert(meesage)
	case LOG_EMERG:
		customlog.Emerg(meesage)
	default:
		customlog.Info(meesage)
	}

}
