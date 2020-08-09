package autoupdate

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
