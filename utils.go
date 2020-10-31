//
// utils.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package utils

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		return
	}
}

// ReadFile() reads a file and returns string content
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
