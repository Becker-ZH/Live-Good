package main

import "database/sql"

// Config ... Configuration parameters for database connection
type Config struct {
	HOST         string `json:"host"`
	USER         string `json:"user"`
	PASSWORD     string `json:"password"`
	DBNAME       string `json:"dbname"`
	JHUAPARTMENT string `json:"jhuapartment"`
	CRIME        string `json:"crime"`
	PREFERENCE   string `json:"preference"`
	PORT         int    `json:"port"`
}

// Request ...
type Request struct {
	Type      string  `json:"type,omitempty"`
	UserName  string  `json:"username"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ApartmentInfo ...
type ApartmentInfo struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Latitude    float64         `json:"latitude"`
	Longitude   float64         `json:"longitude"`
	Address     string          `json:"address"`
	Rating      float64         `json:"rating"`
	Price       sql.NullFloat64 `json:"price"`
	FloorPlan   sql.NullString  `json:"floor_plan"`
	Zipcode     sql.NullInt64   `json:"zip_code"`
	LifeIndex   float64         `json:"life_index"`
	UntilityNum Fake            `json:"utility_num"`
	CrimNum     Fake            `json:"crime_num"`
}

// Response ...
type Response struct {
	Type          string          `json:"type"`
	SavePrefernce bool            `json:"saved_preference"`
	Recommend     []ApartmentInfo `json:"recommend"`
	Status        int             `json:"status_code"`
}

// Fake ...
type Fake struct {
	Int64 int64 `json:"Int64"`
	Valid bool  `json:"Valid"`
}

// Apartments ... for sort
type Apartments []ApartmentInfo

func (a Apartments) Len() int {
	return len(a)
}

func (a Apartments) Less(i, j int) bool {
	return a[i].LifeIndex > a[j].LifeIndex
}

func (a Apartments) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
