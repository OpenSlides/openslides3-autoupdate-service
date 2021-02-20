package vote

import "fmt"

type invalidInputError struct {
	msg string
}

func invalidInput(format string, a ...interface{}) invalidInputError {
	return invalidInputError{fmt.Sprintf(format, a...)}
}

func (e invalidInputError) ClientError() string {
	return "invalid-vote"
}

func (e invalidInputError) Error() string {
	return e.msg
}
