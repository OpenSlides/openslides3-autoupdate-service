package datastore

import "fmt"

type doesNotExist string

func (e doesNotExist) Error() string {
	return fmt.Sprintf("%s does not exist", string(e))
}

func (e doesNotExist) DoesNotExist() {}
