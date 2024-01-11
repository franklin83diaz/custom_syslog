package main

import (
	"github.com/franklin83diaz/custom_syslog/customlog"
)

func main() {
	//see with command tail -f /var/log/syslog
	customlog.Log("log of exaples", customlog.LOG_INFO)

	//this send show menssage in all terminals when you run as root
	customlog.Log("log of exaples LOG_EMERG", customlog.LOG_EMERG)
}
