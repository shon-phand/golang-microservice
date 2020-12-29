package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		"mysqlpasswd",
		"127.0.0.1:3306",
		"bookstore_users",
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = Client.Ping()
	if err != nil {
		fmt.Println("ping error")
		panic(err)
	}

	log.Println("Database successully connected")
}
