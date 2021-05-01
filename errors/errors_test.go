/*
errors_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package errors
package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestUserDefinedError(t *testing.T) {
	var ERR_SOMETHING_WRONG UserDefinedError = "something wrong"
	var ERR_SOMETHING_WRONG_COPY UserDefinedError = ERR_SOMETHING_WRONG

	var err interface{} = ERR_SOMETHING_WRONG
	val, ok := err.(error)
	// is of type `error`
	if !(val.Error() == "something wrong") || !ok {
		t.Error()
	}

	// string value matches
	if !(ERR_SOMETHING_WRONG.Error() == "something wrong") {
		t.Error()
	}

	// new error created using fmt.Errorf() is NOT the same
	if ERR_SOMETHING_WRONG == fmt.Errorf("something wrong") {
		t.Error()
	}

	// new error created using fmt.Errorf() is NOT the same
	if errors.Is(ERR_SOMETHING_WRONG, fmt.Errorf("something wrong")) {
		t.Error()
	}

	// however, the string value of the new error created using fmt.Errorf() is always same
	if !(ERR_SOMETHING_WRONG.Error() == fmt.Errorf("something wrong").Error()) {
		t.Error()
	}

	// new error created using errors.New() is NOT the same
	if errors.Is(ERR_SOMETHING_WRONG, errors.New("something wrong")) {
		t.Error()
	}

	// however, the string value of the new error created using errors.New() is always same
	if !(ERR_SOMETHING_WRONG.Error() == errors.New("something wrong").Error()) {
		t.Error()
	}

	// a copy/reference of two errors are always same
	if !(ERR_SOMETHING_WRONG == ERR_SOMETHING_WRONG_COPY) {
		t.Error()
	}

	// a copy/reference of two errors are always same
	if !errors.Is(ERR_SOMETHING_WRONG, ERR_SOMETHING_WRONG_COPY) {
		t.Error()
	}

	// the string values of a copy/reference of two errors are always same
	if !(ERR_SOMETHING_WRONG.Error() == ERR_SOMETHING_WRONG_COPY.Error()) {
		t.Error()
	}
}
