package main

import (
	"os"
	"strconv"
)

func getConfigDetailes() (Config, error) {
	var dbConn Config

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return dbConn, err
	}

	dbConn = Config{
		HOST:      os.Getenv("HOST"),
		USER:      os.Getenv("USER"),
		PASSWORD:  os.Getenv("PASSWORD"),
		DBNAME:    os.Getenv("DBNAME"),
		TABLENAME: os.Getenv("TABLENAME"),
		PORT:      port,
	}

	return dbConn, nil

}
