package utils

import "testing"

func TestReadFile(t *testing.T) {
	data, err := ReadFile("")
	if data == "" && err != nil {
		//
	} else {
		t.Errorf("")
	}
}
