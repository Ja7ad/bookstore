package services

import (
	"bookstore/domain/users"
	"bookstore/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
