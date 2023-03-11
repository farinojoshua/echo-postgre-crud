package db

import (
	"database/sql"
	"echo-user-app/config"

	_ "github.com/lib/pq" //
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := "postgres://" + conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@" + conf.DB_HOST + "/" + conf.DB_NAME + "?sslmode=disable"

	// "postgres://postgres:123@localhost/stocksdb?sslmode=disable"

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic("ConnectionString Error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}

}

func CreateCon() *sql.DB {
	return db
}
