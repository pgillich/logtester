/*
	Package errorformatter_tester test github.com/pgillich/errorformatter
*/
// nolint:stylecheck,golint
package errorformatter_tester

import (
	"fmt"
	"net/http"
	"os"

	"github.com/juju/rfc/rfc5424"
	"github.com/pgillich/errorformatter"
	log "github.com/sirupsen/logrus"
)

// TryErrorformatter checks the given formatter
// nolint:gocyclo
func TryErrorformatter(formatterName string, testCase string, flags int, callStackSkipLast int) {
	errorformatter.AddSkipPackageFromStackTrace("github.com/pgillich/logtester")

	var logger *log.Logger
	switch formatterName {
	case "text":
		logger = errorformatter.NewTextLogger(log.InfoLevel, flags, callStackSkipLast)
	case "syslog":
		logger = errorformatter.NewSyslogLogger(log.InfoLevel, flags, callStackSkipLast,
			rfc5424.FacilityDaemon, rfc5424.Hostname{FQDN: "fqdn.host.com"}, "application",
			"PID", "")
	case "json":
		logger = errorformatter.NewJSONLogger(log.InfoLevel, flags, callStackSkipLast)
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
	err := errorformatter.GenerateDeepErrors()

	fmt.Println()
	respBody := []byte{}
	errorformatter.ExtractHTTPProblem(&respBody,
		logger.WithError(err), log.ErrorLevel, http.StatusPreconditionFailed,
	).Log(log.ErrorLevel, "USER MSG")
	fmt.Printf("\n%s\n", string(respBody))
}

func tryError(logger *log.Logger) {
	err := errorformatter.GenerateDeepErrors()

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
