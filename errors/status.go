package errors

import (
	"net/http"
)

var (
	_ Option = StatusCode(0)
)

// Known status codes.
const (
	StatusContinue                      = StatusCode(http.StatusContinue)
	StatusSwitchingProtocols            = StatusCode(http.StatusSwitchingProtocols)
	StatusProcessing                    = StatusCode(http.StatusProcessing)
	StatusEarlyHints                    = StatusCode(http.StatusEarlyHints)
	StatusOK                            = StatusCode(http.StatusOK)
	StatusCreated                       = StatusCode(http.StatusCreated)
	StatusAccepted                      = StatusCode(http.StatusAccepted)
	StatusNonAuthoritativeInfo          = StatusCode(http.StatusNonAuthoritativeInfo)
	StatusNoContent                     = StatusCode(http.StatusNoContent)
	StatusResetContent                  = StatusCode(http.StatusResetContent)
	StatusPartialContent                = StatusCode(http.StatusPartialContent)
	StatusMultiStatus                   = StatusCode(http.StatusMultiStatus)
	StatusAlreadyReported               = StatusCode(http.StatusAlreadyReported)
	StatusIMUsed                        = StatusCode(http.StatusIMUsed)
	StatusMultipleChoices               = StatusCode(http.StatusMultipleChoices)
	StatusMovedPermanently              = StatusCode(http.StatusMovedPermanently)
	StatusFound                         = StatusCode(http.StatusFound)
	StatusSeeOther                      = StatusCode(http.StatusSeeOther)
	StatusNotModified                   = StatusCode(http.StatusNotModified)
	StatusUseProxy                      = StatusCode(http.StatusUseProxy)
	StatusTemporaryRedirect             = StatusCode(http.StatusTemporaryRedirect)
	StatusPermanentRedirect             = StatusCode(http.StatusPermanentRedirect)
	StatusBadRequest                    = StatusCode(http.StatusBadRequest)
	StatusUnauthorized                  = StatusCode(http.StatusUnauthorized)
	StatusPaymentRequired               = StatusCode(http.StatusPaymentRequired)
	StatusForbidden                     = StatusCode(http.StatusForbidden)
	StatusNotFound                      = StatusCode(http.StatusNotFound)
	StatusMethodNotAllowed              = StatusCode(http.StatusMethodNotAllowed)
	StatusNotAcceptable                 = StatusCode(http.StatusNotAcceptable)
	StatusProxyAuthRequired             = StatusCode(http.StatusProxyAuthRequired)
	StatusRequestTimeout                = StatusCode(http.StatusRequestTimeout)
	StatusConflict                      = StatusCode(http.StatusConflict)
	StatusGone                          = StatusCode(http.StatusGone)
	StatusLengthRequired                = StatusCode(http.StatusLengthRequired)
	StatusPreconditionFailed            = StatusCode(http.StatusPreconditionFailed)
	StatusRequestEntityTooLarge         = StatusCode(http.StatusRequestEntityTooLarge)
	StatusRequestURITooLong             = StatusCode(http.StatusRequestURITooLong)
	StatusUnsupportedMediaType          = StatusCode(http.StatusUnsupportedMediaType)
	StatusRequestedRangeNotSatisfiable  = StatusCode(http.StatusRequestedRangeNotSatisfiable)
	StatusExpectationFailed             = StatusCode(http.StatusExpectationFailed)
	StatusTeapot                        = StatusCode(http.StatusTeapot)
	StatusMisdirectedRequest            = StatusCode(http.StatusMisdirectedRequest)
	StatusUnprocessableEntity           = StatusCode(http.StatusUnprocessableEntity)
	StatusLocked                        = StatusCode(http.StatusLocked)
	StatusFailedDependency              = StatusCode(http.StatusFailedDependency)
	StatusTooEarly                      = StatusCode(http.StatusTooEarly)
	StatusUpgradeRequired               = StatusCode(http.StatusUpgradeRequired)
	StatusPreconditionRequired          = StatusCode(http.StatusPreconditionRequired)
	StatusTooManyRequests               = StatusCode(http.StatusTooManyRequests)
	StatusRequestHeaderFieldsTooLarge   = StatusCode(http.StatusRequestHeaderFieldsTooLarge)
	StatusUnavailableForLegalReasons    = StatusCode(http.StatusUnavailableForLegalReasons)
	StatusInternalServerError           = StatusCode(http.StatusInternalServerError)
	StatusNotImplemented                = StatusCode(http.StatusNotImplemented)
	StatusBadGateway                    = StatusCode(http.StatusBadGateway)
	StatusServiceUnavailable            = StatusCode(http.StatusServiceUnavailable)
	StatusGatewayTimeout                = StatusCode(http.StatusGatewayTimeout)
	StatusHTTPVersionNotSupported       = StatusCode(http.StatusHTTPVersionNotSupported)
	StatusVariantAlsoNegotiates         = StatusCode(http.StatusVariantAlsoNegotiates)
	StatusInsufficientStorage           = StatusCode(http.StatusInsufficientStorage)
	StatusLoopDetected                  = StatusCode(http.StatusLoopDetected)
	StatusNotExtended                   = StatusCode(http.StatusNotExtended)
	StatusNetworkAuthenticationRequired = StatusCode(http.StatusNetworkAuthenticationRequired)
)

// StatusCode describes an HTTP status code.
type StatusCode int

// Int returns the status code as "int".
func (s StatusCode) Int() int {
	return int(s)
}

// String implements the fmt.Stringer interface.
func (s StatusCode) String() string {
	return http.StatusText(s.Int())
}

// In returns true if the given error has this status code.
func (s StatusCode) In(err error) bool {
	return s == GetStatusCode(err)
}

// Apply implements the Option interface.
func (s StatusCode) Apply(_ bool, err error) {
	if e, ok := err.(*wrappedError); ok {
		e.statusCode = s
	}
}

// GetStatusCode gets the status code from the error, or 0 if not set.
func GetStatusCode(err error) StatusCode {
	if e, ok := err.(*wrappedError); ok {
		return e.statusCode
	}
	return 0
}

// GetStatusCode gets the status code from the error, or InternalServerError if not set.
func GetStatusCodeOr500(err error) StatusCode {
	if e, ok := err.(*wrappedError); ok && e.statusCode != 0 {
		return e.statusCode
	}
	return StatusInternalServerError
}
