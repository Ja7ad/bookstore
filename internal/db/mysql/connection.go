// Package mysql provides a MySQL implementation of the database interface
package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	MYSQL_USERNAME = "mysql_username"
	MYSQL_PASSWORD = "mysql_password"
	MYSQL_HOST     = "mysql_host"
	MYSQL_SCHEMA   = "mysql_schema"
)

var (
	Client *sql.DB // Client is the MySQL connection instance

	username = os.Getenv(MYSQL_USERNAME)
	password = os.Getenv(MYSQL_PASSWORD)
	host     = os.Getenv(MYSQL_HOST)
	schema   = os.Getenv(MYSQL_SCHEMA)
)

// init registers the MySQL driver
func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configurated")
}
