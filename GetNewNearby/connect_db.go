package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func connectDB(dbConn Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConn.HOST, dbConn.PORT, dbConn.USER, dbConn.PASSWORD, dbConn.DBNAME)
	db, err := sql.Open("postgres", psqlInfo)

	return db, err
}
