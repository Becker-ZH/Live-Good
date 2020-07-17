package main

import (
	"math"
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
		CRIME:        os.Getenv("CRIME"),
		PREFERENCE:   os.Getenv("PREFERENCE"),
		PORT:         port,
	}

	return dbConn, nil

}

func earthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137
	//define radius of earth

	rad := math.Pi / 180.0
	//define radian

	//calculate distance
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius * 0.621371
}
