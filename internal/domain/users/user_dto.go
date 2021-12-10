// Package users data transfer object
package users

import (
	"bookstore/pkg/errors/restError"
	"strings"
)

const (
	StatusActive = "active"
)

// User entity
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

// Validate user entity
func (u *User) Validate() *restError.RestErr {
	u.FirstName = strings.TrimSpace(strings.ToLower(u.FirstName))
	u.LastName = strings.TrimSpace(strings.ToLower(u.LastName))

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return restError.NewBadRequestError("invalid email address")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return restError.NewBadRequestError("invalid user password")
	}

	return nil
}
