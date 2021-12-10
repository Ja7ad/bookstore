// Package users Data Access Object
package users

import (
	"bookstore/internal/db/msql"
	"bookstore/pkg/errors/restError"
	"bookstore/pkg/logger"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Mysql queries
const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser    = " id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUser   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// Get user by id from database
func (u *User) Get() *restError.RestErr {
	stmt, err := msql.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return restError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return restError.NewInternalServerError("database error")
	}

	return nil
}

// Save user to database
func (u *User) Save() *restError.RestErr {
	stmt, err := msql.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return restError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// insetResult exec query for save user to database
	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user statement", err)
		return restError.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return restError.NewInternalServerError("database error")
	}

	u.Id = userId
	return nil
}

// Update user in database
func (u *User) Update() *restError.RestErr {
	stmt, err := msql.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return restError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return restError.NewInternalServerError("database error")
	}

	return nil

}

// Delete user in database
func (u *User) Delete() *restError.RestErr {
	stmt, err := msql.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return restError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return restError.NewInternalServerError("database error")
	}

	return nil
}

// FindByStatus in database by status
func (u *User) FindByStatus(status string) ([]User, *restError.RestErr) {
	stmt, err := msql.Client.Prepare(queryFindUser)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, restError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, restError.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan row into user structure", err)
			return nil, restError.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, restError.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
