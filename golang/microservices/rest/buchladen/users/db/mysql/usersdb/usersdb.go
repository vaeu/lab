package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_usersdb_name   = "mysql_name"
	mysql_usersdb_pass   = "mysql_pass"
	mysql_usersdb_host   = "mysql_host"
	mysql_usersdb_schema = "mysql_schema"
)

var (
	Client *sql.DB
	name   = os.Getenv(mysql_usersdb_name)
	pass   = os.Getenv(mysql_usersdb_pass)
	host   = os.Getenv(mysql_usersdb_host)
	schema = os.Getenv(mysql_usersdb_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		name, pass, host, schema,
	)
	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("db has been successfully configured")
}
