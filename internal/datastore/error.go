package datastore

import (
	"errors"
	"fmt"
	"strings"
)

type doesNotExistError string

func (e doesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist", string(e))
}

func (e doesNotExistError) DoesNotExist() string {
	return string(e)
}

type resetError struct{}

func (e resetError) Error() string {
	return "reset needed"
}

func (e resetError) Reset() {}

type conditionError struct {
	condition *DebugCondition
	err       error
}

func (e conditionError) Error() string {
	return e.err.Error()
}

func (e conditionError) Unwrap() error {
	return e.err
}

func (e conditionError) Conditions() []string {
	con := e.condition.getConditions()
	var sErr conditionError
	if errors.As(e.err, &sErr) {
		con = append(con, sErr.Conditions()...)
	}
	return con
}

func (e conditionError) ConditionError() error {
	var dErr interface {
		Error() string
		DoesNotExist() string
	}
	if !errors.As(e, &dErr) {
		return fmt.Errorf("invalid condition error: conditions without doesNotExist")
	}

	return fmt.Errorf(
		"Database is invalid. Could not find element %s.\nThe element is required under the following conditions:\n%s\nError: %s",
		dErr.DoesNotExist(),
		strings.Join(e.Conditions(), "\n"),
		e.Error(),
	)
}

// DebugCondition is used to get informations, in which constellation the database
// was invalid.
//
// To do this, create a DebugCondition at the beginning of a function and after each
// state-ceck, append the condition of the database.
//
// If an error happens, that could be a doesNotExistError, then wrap the error
// with DebugCondition.Error().
//
// Example:
//
//	func foo(ds DataStore, u User) error{
//	con := new(datastore.DebugCondition)
//	if u.IsPresent() {
//		return nil
//	}
//	conn.Append("User is not present")
//
//	var username struct {
//		Username string
//	}
//	if err := ds.Get("users/user:"+u.ID, &username); err != nil {
//		return conn.Error("getting username: %w", err)
//	}
//	return nil
//	}
type DebugCondition struct {
	conditions []string
	parent     *DebugCondition
}

// Append adds one conditions description
func (c *DebugCondition) Append(format string, a ...interface{}) {
	c.conditions = append(c.conditions, fmt.Sprintf(format, a...))
}

// Error creates an error under the condition.
func (c *DebugCondition) Error(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)

	var sErr *conditionError
	var dErr interface {
		DoesNotExist() string
	}
	if errors.As(err, &sErr) || errors.As(err, &dErr) {
		return conditionError{
			condition: c,
			err:       err,
		}
	}

	return err
}

// Sub creates a new Condition that contains all conditions of the given
// condition.
func (c *DebugCondition) Sub() *DebugCondition {
	return &DebugCondition{parent: c}
}

func (c *DebugCondition) getConditions() []string {
	if c.parent == nil {
		return c.conditions
	}

	return append(c.parent.getConditions(), c.conditions...)
}
