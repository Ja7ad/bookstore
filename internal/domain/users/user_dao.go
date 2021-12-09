// Package users Data Access Object
package users

import (
	"bookstore/errors/mysqlError"
	"bookstore/errors/restError"
	"bookstore/internal/db/msql"
	"bookstore/utils/date"
	_ "github.com/go-sql-driver/mysql"
)

// Mysql queries
const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get user by id from database
func (u *User) Get() *restError.RestErr {
	stmt, err := msql.Client.Prepare(queryGetUser)
	if err != nil {
		return restError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return mysqlError.ParseError(getErr)
	}

	return nil
}

// Save user to database
func (u *User) Save() *restError.RestErr {
	stmt, err := msql.Client.Prepare(queryInsertUser)
	if err != nil {
		return restError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	u.DateCreated = date.GetNowString()

	// insetResult exec query for save user to database
	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, date.GetNowString())
	if saveErr != nil {
		return mysqlError.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlError.ParseError(err)
	}

	u.Id = userId
	return nil
}

// Update user in database
func (u *User) Update() *restError.RestErr {

}
