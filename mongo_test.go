package main

import (
	"testing"
	"time"
)

func Test_getImage(t *testing.T) {
	_, err := getImages("uhajzeraj")

	if err != nil {
		t.Error(err)
	}
}
func Test_deleteUser(t *testing.T) {
	err := deleteUser("testUser")

	if err == nil {
		t.Error(err)
	}
}
func Test_changeUserPassword(t *testing.T) {
	err := changeUserPassword("testUser", "test123")
	if err == nil {
		t.Error(err)
	}
}
func Test_storeImage(t *testing.T) {
	err := storeImage("testUser", "test123", nil)
	if err == nil {
		t.Error(err)
	}
}
func Test_removeImage(t *testing.T) {
	err := removeImage("uhajzeraj", "test")

	if err != nil {
		t.Error(err)
	}
}
func Test_findEmail(t *testing.T) {
	_, err := findEmail("uhajzeraj")

	if err != nil {
		t.Error(err)
	}
}

func Test_storeFailedLogin(t *testing.T) {
	err := storeFailedLogin("uhajzeraj", time.Now())

	if err != nil {
		t.Error(err)
	}
}

func Test_countLogins(t *testing.T) {
	i := countLogins("uhajzeraj")

	if i < 0 {
		t.Error("CountLogin is not working properly!")
	}
}
