package main

import (
	"testing"
)

func Test_errorPanic(t *testing.T) {
	var err error
	errorPanic(err)
}
func Test_exists(t *testing.T) {
	boolValue, err := exists("../stegano/.idea")
	if err != nil {
		t.Error("Error:Doesnt exist")
	}
	if boolValue != true {
		t.Error("Path dont exist")
	}
}

func Test_validateSignup(t *testing.T) {

	err := validateSignup("etnik5", "etnik123", "etnik123", "etnikg@stud.ntnu.no")
	if err != nil {
		t.Errorf("Error validating Signup:%s", err)
	}
}

func Test_validateLogin(t *testing.T) {

	err := validateLogin("etnik5", "etnik123")
	if err != nil {
		t.Errorf("Error validating Login:%s", err)
	}
}
