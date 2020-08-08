package datastore

import "fmt"

type doesNotExistError string

func (e doesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist", string(e))
}

func (e doesNotExistError) DoesNotExist() {}

type resetError struct{}

func (e resetError) Error() string {
	return "reset needed"
}

func (e resetError) Reset() {}
