package autoupdate

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
	return fmt.Sprintf("Invalid request: %v", e.err)
}

func (e invalidRequestError) ClientError() string {
	return "invalid_request"
}
