package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Config ... Configuration parameters for database connection
type Config struct {
	HOST      string `json:"host"`
	USER      string `json:"user"`
	PASSWORD  string `json:"password"`
	DBNAME    string `json:"dbname"`
	TABLENAME string `json:"tablename"`
	PORT      int    `json:"port"`
}

// Coordinates ... Coordinates information for apartment
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ApartmentInfo ... Apartment information
type ApartmentInfo struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"desc"`
	Coordinate  Coordinates `json:"coordinates"`
}

// AllApartment ... return all apartment information
type AllApartment struct {
	AllApartmentInfo []ApartmentInfo `json:"apartments"`
}

func handler() (AllApartment, error) {

	var response AllApartment

	dbConn, err := getConfigDetailes()
	if err != nil {
		return response, err
	}

	response, err = readAll(dbConn)
	return response, err

}

func main() {
	lambda.Start(handler)
}
