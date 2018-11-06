package main

import (
	"testing"
)

func Test_errorPanic(t *testing.T) {
	var err error
	errorPanic(err)
}
