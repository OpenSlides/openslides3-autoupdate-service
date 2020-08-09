package autoupdate

type invalidRequestError struct {
	msg string
}

func (e invalidRequestError) Error() string {
	return e.msg
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
