package services

import (
	"bookstore/errors/restError"
	"bookstore/internal/domain/users"
)

// GetUser gets a user by its id
func GetUser(userId int64) (*users.User, *restError.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser creates a new user
func CreateUser(user users.User) (*users.User, *restError.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user
func UpdateUser(user users.User) (*users.User, *restError.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	current.Update()
}
