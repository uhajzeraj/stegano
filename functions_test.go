package main

import (
	"testing"
)

func Test_errorPanic(t *testing.T) {
	var err error
	errorPanic(err)
}
func Test_returnEmptyError(t *testing.T) {
	var err error
	errorPanic(err)
}
