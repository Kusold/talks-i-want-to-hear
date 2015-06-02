package models

import "testing"

func TestCreateUser(t *testing.T) {
	user := User{
		0,
		"test@example.com",
		"foo",
	}
	err := user.CreateUser()
	if err != nil {
		t.Fatal(err.Error())
	}
}
