/*
errors.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package errors holds some error related utils.
package errors

type UserDefinedError string

func (e UserDefinedError) Error() string {
	return string(e)
}
