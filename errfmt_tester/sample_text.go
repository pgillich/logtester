// nolint:stylecheck,golint
package errfmt_tester

import (
	"github.com/pgillich/errfmt"
	log "github.com/sirupsen/logrus"
)

func trySampleText() {
	// register a trim prefix (optional)
	errfmt.AddSkipPackageFromStackTrace("github.com/pgillich/logtester")

	// build a new logrus logger
	logger := errfmt.NewTextLogger(log.InfoLevel, errfmt.FlagNone, 0)

	// Info log with key-value map
	logger.WithFields(log.Fields{
		"STR":  "str",
		"INT":  42,
		"BOOL": true,
	}).Info("USER MSG")
}
