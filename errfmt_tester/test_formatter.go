/*
	Package errfmt_tester test github.com/pgillich/errfmt
*/
// nolint:stylecheck,golint
package errfmt_tester

import (
	"fmt"
	"net/http"
	"os"

	"github.com/juju/rfc/rfc5424"
	"github.com/pgillich/errfmt"
	log "github.com/sirupsen/logrus"
)

// TryErrorformatter checks the given formatter
// nolint:gocyclo
func TryErrorformatter(formatterName string, testCase string, flags int, callStackSkipLast int) {
	errfmt.AddSkipPackageFromStackTrace("github.com/pgillich/logtester")

	var logger *log.Logger
	switch formatterName {
	case "text":
		logger = errfmt.NewTextLogger(log.InfoLevel, flags, callStackSkipLast)
	case "syslog":
		logger = errfmt.NewSyslogLogger(log.InfoLevel, flags, callStackSkipLast,
			rfc5424.FacilityDaemon, rfc5424.Hostname{FQDN: "fqdn.host.com"}, "application",
			"PID", "")
	case "json":
		logger = errfmt.NewJSONLogger(log.InfoLevel, flags, callStackSkipLast)
	default:
		fmt.Printf("Unknown formatter: %s\n", formatterName)
		os.Exit(2)
	}

	switch testCase {
	case "sampletext":
		trySampleText()
	case "samplehttp":
		trySampleHTTP()
	case "info":
		tryInfo(logger)
	case "error":
		tryError(logger)
	case "errorhttp":
		tryErrorHTTP(logger)
	default:
		fmt.Printf("Unknown test case: %s\n", testCase)
		os.Exit(3)
	}
}

func tryErrorHTTP(logger *log.Logger) {
	err := errfmt.GenerateDeepErrors()

	fmt.Println()
	respBody := []byte{}
	errfmt.ExtractHTTPProblem(&respBody, http.StatusPreconditionFailed,
		logger.WithError(err)).Error("USER MSG")
	fmt.Printf("\n%s\n", string(respBody))
}

func tryError(logger *log.Logger) {
	err := errfmt.GenerateDeepErrors()

	fmt.Println()
	logger.WithError(err).Error("USER MSG")
}

func tryInfo(logger *log.Logger) {
	logger.WithFields(log.Fields{
		"STR":  "str",
		"INT":  42,
		"BOOL": true,
	}).Info("USER MSG")
}
