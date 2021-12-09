// Package users data transfer object
package users

import (
	"bookstore/errors/restError"
	"strings"
)

// User entity
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate user entity
func (u *User) Validate() *restError.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return restError.NewBadRequestError("invalid email address")
	}

	return nil
}
