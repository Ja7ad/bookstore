package services

import (
	"bookstore/internal/domain/users"
	"bookstore/pkg/crypto_tool"
	"bookstore/pkg/date"
	"bookstore/pkg/errors/restError"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *restError.RestErr)
	CreateUser(users.User) (*users.User, *restError.RestErr)
	UpdateUser(bool, users.User) (*users.User, *restError.RestErr)
	DeleteUser(int64) *restError.RestErr
	SearchUser(string) (users.Users, *restError.RestErr)
}

// GetUser gets a user by its id
func (s *usersService) GetUser(userId int64) (*users.User, *restError.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser creates a new user
func (s *usersService) CreateUser(user users.User) (*users.User, *restError.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBFormat()
	user.Password = crypto_tool.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *restError.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
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
func (s *usersService) DeleteUser(userId int64) *restError.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

// SearchUser user in database
func (s *usersService) SearchUser(status string) (users.Users, *restError.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
