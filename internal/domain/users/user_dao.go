// Package users Data Access Object
package users

import (
	"bookstore/internal/db/mysql"
	"bookstore/utils/date"
	"bookstore/utils/errors"
	"fmt"
	"strings"
)

const (
	indexUniqueEmail = "users_email_uindex"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User) // usersDB is a map of users
)

// Get user by id from database
func (u *User) Get() *errors.RestErr {
	if err := mysql.Client.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	u.DateCreated = date.GetNowString()

	// insetResult exec query for save user to database
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, date.GetNowString())
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", u.Email))
		}
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	u.Id = userId
	return nil
}
