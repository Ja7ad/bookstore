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
func UpdateUser(isPartial bool, user users.User) (*users.User, *restError.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if errUpd := current.Update(); errUpd != nil {
		return nil, errUpd
	}

	return current, nil
}

// DeleteUser delete a user
func DeleteUser(userId int64) *restError.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}
