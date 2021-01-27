package http

import "fmt"

type invalidRequestError struct {
	err error
}

func (e invalidRequestError) Error() string {
	return fmt.Sprintf("Invalid request: %v", e.err)
}

func (e invalidRequestError) ClientError() string {
	return "invalid_request"
}

// noStatusCodeError helps the errorHandler do decide, if an status code can be
// set.
type noStatusCodeError struct {
	wrapped error
}

func (e noStatusCodeError) Error() string {
	return e.wrapped.Error()
}

func (e noStatusCodeError) Unwrap() error {
	return e.wrapped
}

func (e noStatusCodeError) NoStatus() {}

type authRequiredError struct {
	msg string
}

func (e authRequiredError) Error() string {
	msg := e.msg
	if msg == "" {
		msg = "Access for anonymous user is disabled"
	}
	return msg
}

func (e authRequiredError) ClientError() string {
	return "auth_required"
}
