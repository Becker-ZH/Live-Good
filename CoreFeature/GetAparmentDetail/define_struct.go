package main

import "database/sql"

// Config ... Configuration parameters for database connection
type Config struct {
	HOST      string `json:"host"`
	USER      string `json:"user"`
	PASSWORD  string `json:"password"`
	DBNAME    string `json:"dbname"`
	TABLENAME string `json:"table"`
	PORT      int    `json:"port"`
}

// Request ...
type Request struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Crime ...
type Crime struct {
	Neighborhood string  `json:"neighborhood"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Location     string  `json:"location"`
}

// Shop ...
type Shop struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// ApartmentInfo ...
type ApartmentInfo struct {
	Address    string          `json:"address"`
	Rating     float64         `json:"rating"`
	Price      sql.NullFloat64 `json:"price"`
	FloorPlan  sql.NullString  `json:"floor_plan"`
	Restaurant sql.NullInt64   `json:"restaurant_number"`
	Market     sql.NullInt64   `json:"market_number"`
	Gas        sql.NullInt64   `json:"gas_number"`
	Bar        sql.NullInt64   `json:"bar"`
	Park       sql.NullInt64   `json:"park"`
	Gym        sql.NullInt64   `json:"gym"`
	Hotel      sql.NullInt64   `json:"hotel"`
	Zipcode    sql.NullInt64   `json:"zip_code"`
	CrimeNum   Fake            `json:"crime_num"`
	UtilityNum Fake            `json:"utility_num"`
}

// Response ...
type Response struct {
	Name      string        `json:"name"`
	Longitude float64       `json:"longitude"`
	Latitude  float64       `json:"latitude"`
	Apartment ApartmentInfo `json:"apartment_info"`
	CrimeData []Crime       `json:"crime_data"`
}

// Fake ...
type Fake struct {
	Int64 int64 `json:"Int64"`
	Valid bool  `json:"Valid"`
}
