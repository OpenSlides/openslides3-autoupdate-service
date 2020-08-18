package projector

import "fmt"

// ClientError is a error message for the client.
type ClientError struct {
	msg string
}

// NewClientError creates a new ClientError with a formatted message.
func NewClientError(format string, a ...interface{}) ClientError {
	return ClientError{fmt.Sprintf(format, a...)}
}

func (c ClientError) Error() string {
	return c.msg
}
