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

func TestHasValidCredentials(t *testing.T) {
	user := User{
		0,
		"test-login@example.com",
		"foo",
	}
	err := user.CreateUser()
	if err != nil {
		t.Fatal(err.Error())
	}

	userResult, err := user.HasValidCredentials()
	if err != nil {
		t.Fatal(err)
	}

	if userResult.Email != user.Email {
		t.Fatal("Emails don't match")
	}

	if userResult.Password != user.Password {
		t.Fatal("Emails don't match")
	}
}

func TestHasValidCredentialsFailureCase(t *testing.T) {
	user := User{
		0,
		"test-doesntexist@example.com",
		"foo",
	}

	_, err := user.HasValidCredentials()
	if err == nil {
		t.Fatal("Somehow we found a user that we didn't expect to exist")
	}
}
