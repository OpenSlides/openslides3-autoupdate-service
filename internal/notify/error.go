package notify

import "fmt"

type authRequiredError struct{}

func (e authRequiredError) Error() string {
	return "Only authenticated users can use the notify system."
}

func (e authRequiredError) ClientError() string {
	return "auth_required"
}

type invalidRequestError struct {
	err error
}

func (e invalidRequestError) Error() string {
	return fmt.Sprintf("The request body is invalid: %v", e.err)
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

func (e noStatusCodeError) NoStatus() {}
