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
	MysqlUsername = "mysql_username"
	MysqlPassword = "mysql_password"
	MysqlHost     = "mysql_host"
	MysqlSchema   = "mysql_schema"
)

var (
	Client *sql.DB // Client is the MySQL connection instance

	username = os.Getenv(MysqlUsername)
	password = os.Getenv(MysqlPassword)
	host     = os.Getenv(MysqlHost)
	schema   = os.Getenv(MysqlSchema)
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
