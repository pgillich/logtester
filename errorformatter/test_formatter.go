/*
	Package errorformatter test github.com/pgillich/errorformatter
*/
package errorformatter

import (
	"fmt"
	"net/http"
	"os"

	"github.com/juju/rfc/rfc5424"
	"github.com/pgillich/errorformatter"
	log "github.com/sirupsen/logrus"
)

// TryErrorformatter checks the given formatter
func TryErrorformatter(formatterName string, flags int, callStackSkipLast int) {
	var logger *log.Logger
	switch formatterName {
	case "text":
		logger = errorformatter.NewTextLogger(log.InfoLevel, flags, callStackSkipLast)
	case "syslog":
		logger = errorformatter.NewSyslogLogger(log.InfoLevel, flags, callStackSkipLast,
			rfc5424.FacilityDaemon, rfc5424.Hostname{FQDN: "fqdn.host.com"}, "application",
			"PID", "")
	default:
		fmt.Printf("Unknown formatter: %s\n", formatterName)
		os.Exit(2)
	}
	logger.Out = os.Stdout

	err := errorformatter.GenerateDeepErrors()

	fmt.Println()
	//logger.WithError(err).Log(log.ErrorLevel, "USER MSG")
	respBody := []byte{}
	errorformatter.ExtractHTTPProblem(&respBody,
		logger.WithError(err), log.ErrorLevel, http.StatusPreconditionFailed,
	).Log(log.ErrorLevel, "USER MSG")
	fmt.Printf("\n%s\n", string(respBody))
}
