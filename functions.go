package main

import (
	"os"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// UserInfo structure to hold the fetched user hashed pass
type UserInfo struct {
	User     string `bson:"user"`
	HashPass string `bson:"passHash"`
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

// Validate signup
func validateSignup(user, pass, passConfirm, email string) []string {
	var errorSlice []string

	// Check if username is OK
	match, err := regexp.MatchString(`^[a-zA-Z0-9_-]{6,30}$`, user)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Username does not meet the requirements")
	} else {
		// Check if username exists
		exists, err2 := entryExists("user", user, "users")
		returnEmptyError(err2)
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
		exists, err3 := entryExists("email", email, "users")
		returnEmptyError(err3)
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

// Validate login
func validateLogin(user, pass string) []string {
	var errorSlice []string

	// Check if username is OK
	match, err := regexp.MatchString(`^[a-zA-Z0-9_-]{6,30}$`, user)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Username does not meet the requirements")
	}

	// Check if password is OK
	match, err = regexp.MatchString(`^.{6,40}$`, pass)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Password does not meet the requirements")
	}

	// Continue further if there are no errors
	if len(errorSlice) == 0 {

		result := UserInfo{}

		conn.DB("stegano").C("users").Find(bson.M{"user": user}).Select(bson.M{"user": user, "passHash": 1}).One(&result)

		err := bcrypt.CompareHashAndPassword([]byte(result.HashPass), []byte(pass))
		if err != nil {
			errorSlice = append(errorSlice, "Username or password is not correct")
		}
	}

	return errorSlice
}

func validateChangePassword(current, new, confirm, user string) []string {
	var errorSlice []string

	// Check if password is OK
	match, err := regexp.MatchString(`^.{6,40}$`, current)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Old Password does not meet the requirements")
	}

	// Check if newPassword is OK
	match, err = regexp.MatchString(`^.{6,40}$`, new)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "New password does not meet the requirements")
	}

	// Check if confrimPassword is OK
	match, err = regexp.MatchString(`^.{6,40}$`, confirm)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Confirmation password does not meet the requirements")
	}

	// Check if passwords are the same
	if new != confirm {
		errorSlice = append(errorSlice, "Passwords do not match")
	}

	// Continue further if there are no errors
	if len(errorSlice) == 0 {

		result := UserInfo{}

		conn.DB("stegano").C("users").Find(bson.M{"user": user}).Select(bson.M{"user": user, "passHash": 1}).One(&result)

		err := bcrypt.CompareHashAndPassword([]byte(result.HashPass), []byte(current))
		if err != nil {
			errorSlice = append(errorSlice, "Old password is not correct")
		}
	}

	return errorSlice
}
