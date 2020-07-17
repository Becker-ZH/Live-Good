package main

import (
	"os"
	"strconv"
)

func getApartmentConfigDetailes() (Config, error) {
	var dbConn Config

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return dbConn, err
	}

	dbConn = Config{
		HOST:         os.Getenv("HOST"),
		USER:         os.Getenv("USER"),
		PASSWORD:     os.Getenv("PASSWORD"),
		DBNAME:       os.Getenv("DBNAME"),
		JHUAPARTMENT: os.Getenv("JHUAPARTMENT"),
		JHUPLACE:     os.Getenv("JHUPLACE"),
		APTPLACE:     os.Getenv("APTPLACE"),
		PORT:         port,
	}

	return dbConn, nil

}
