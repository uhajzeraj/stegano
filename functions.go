package main

import (
	"os"
	"regexp"
)

// Damn you gometalinter and your cyclomatic complexity
func errorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// Metalinter
func returnEmptyError(err error) {
	if err != nil {
		return
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Validate fields
func validateSignup(user, pass, passConfirm, email string) []string {
	var errorSlice []string

	// Check if username is OK
	match, err := regexp.MatchString(`^[a-zA-Z0-9_-]{6,30}$`, user)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Username does not meet the requirements")
	} else {
		// Check if username exists
		exists, err := entryExists("user", user, "users")
		returnEmptyError(err)
		if exists {
			errorSlice = append(errorSlice, "Username already exists")
		}
	}

	// Check if email is OK
	match, err = regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Email does not meet the requirements")
	} else {
		// Check if email exists
		exists, err := entryExists("email", email, "users")
		returnEmptyError(err)
		if exists {
			errorSlice = append(errorSlice, "Email already exists")
		}
	}

	// Check if password is OK
	match, err = regexp.MatchString(`^.{6,40}$`, pass)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Password does not meet the requirements")
	}

	// Check if confrimPassword is OK
	match, err = regexp.MatchString(`^.{6,40}$`, passConfirm)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Confirmation password does not meet the requirements")
	}

	// Check if passwords are the same
	if pass != passConfirm {
		errorSlice = append(errorSlice, "Passwords do not match")
	}

	return errorSlice
}
