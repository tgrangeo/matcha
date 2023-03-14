package models

import (
    "database/sql"
    "errors"
    _ "github.com/lib/pq"
)


type User struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Type  Type    `json:"type"`
}

func (usr *User) getUser(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (usr *User) updateUser(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (user *User) deleteUser(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (user *User) createUser(db *sql.DB) error {
  return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]User, error) {
  return nil, errors.New("Not implemented")
}