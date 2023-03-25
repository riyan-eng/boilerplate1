package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var PostgreSQLDB *sql.DB

func ConnDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=riyan dbname=boilerplate1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	PostgreSQLDB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("database: can't connect to database")
		os.Exit(1)
	}

	PostgreSQLDB.SetMaxIdleConns(10)
	PostgreSQLDB.SetMaxOpenConns(100)
	PostgreSQLDB.SetConnMaxLifetime(time.Hour)
	if err := PostgreSQLDB.Ping(); err != nil {
		log.Fatal("database: can't ping to database")
	}

	fmt.Println("database: connection opened to database")
}
