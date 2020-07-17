package main

import (
	"os"
	"strconv"
)

func getSchoolConfigDetailes() (Config, error) {
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
		TABLENAME: os.Getenv("SCHOOLTABLE"),
		PORT:      port,
	}

	return dbConn, nil

}

func getApartmentConfigDetailes() (Config, error) {
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
		TABLENAME: os.Getenv("APARTMENTABLE"),
		PORT:      port,
	}

	return dbConn, nil

}
