// nolint:stylecheck,golint
package errorformatter_tester

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/pgillich/errorformatter"
	log "github.com/sirupsen/logrus"
)

func trySampleHTTP() {
	// register a trim prefix (optional)
	errorformatter.AddSkipPackageFromStackTrace("github.com/pgillich/logtester")

	// build a new logrus logger
	logger := errorformatter.NewTextLogger(log.InfoLevel, 0, 0)

	// this decorator sets body, header and status, if response error is NOT nil
	handler := func(w http.ResponseWriter, r *http.Request) {
		if statusCode, err := doRequest(w, r); err != nil { // calling worker func
			errorformatter.WriteHTTPProblem( // sending HTTP error
				w, logger.WithError(err), log.ErrorLevel, statusCode,
			).Log(log.ErrorLevel, "USER MSG") // logging to the console
		}
	}

	// prepare fake request/response
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	// call the decorated handler
	handler(w, req)

	// print out response
	resp := w.Result()
	defer resp.Body.Close()              // nolint:errcheck
	body, _ := ioutil.ReadAll(resp.Body) // nolint:errcheck
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}

/*
doRequest makes the main part of the request.
	if the returned error is nil, the response body, header and status is set
	if the returned error is NOT nil, the response body and status is NOT set (the caller must do it)
*/
// nolint:unparam
func doRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusPreconditionFailed, errorformatter.GenerateDeepErrors()
}
