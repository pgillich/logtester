// nolint:stylecheck,golint
package errorformatter_tester

import (
	"github.com/pgillich/errorformatter"
	log "github.com/sirupsen/logrus"
)

func trySampleText() {
	// register a trim prefix (optional)
	errorformatter.AddSkipPackageFromStackTrace("github.com/pgillich/logtester")

	// build a new logrus logger
	logger := errorformatter.NewTextLogger(log.InfoLevel, errorformatter.FlagNone, 0)

	// Info log with key-value map
	logger.WithFields(log.Fields{
		"STR":  "str",
		"INT":  42,
		"BOOL": true,
	}).Info("USER MSG")
}
