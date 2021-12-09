// Package mysqlError implements the mysqlError error codes.
package mysqlError

import (
	"bookstore/errors/restError"
	"github.com/go-sql-driver/mysql"
	"strings"
)

// Mysql errors
const (
	errorNoRows = "no rows in result set"
)

// ParseError is a wrapper for the mysqlError error
func ParseError(err error) *restError.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return restError.NewNotFoundError("no record matching given id")
		}
		return restError.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return restError.NewBadRequestError("invalid data")
	}

	return restError.NewInternalServerError("error processing request")
}
