package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func (u User) CreateUser() (err error) {
	db, err := sql.Open("postgres", "user=postgres dbname=golang sslmode=disable")
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users(email, password) VALUES($1,$2)", u.Email, u.Password)
	return
}
