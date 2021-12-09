// Package users Data Access Object
package users

import (
	"bookstore/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User) // usersDB is a map of users
)

// Get user by id from database
func (u *User) Get() *errors.RestErr {
	result := usersDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}

	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated

	return nil
}

// Save user to database
func (u *User) Save() *errors.RestErr {
	current := usersDB[u.Id]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.Id))
	}

	usersDB[u.Id] = u

	return nil
}
