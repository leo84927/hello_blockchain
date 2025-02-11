package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"hello_blockchain/config"
	"strings"
)

func main() {

	var (
		dsn      strings.Builder
		dbConfig = config.MysqlConn
	)

	// connection init
	_, err := fmt.Fprintf(&dsn, "%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", dsn.String())
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	// migration start
	m, err := migrate.NewWithDatabaseInstance("file://migration", "mysql", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		panic(err)
	}
}
