package main

// Config ... Configuration parameters for database connection
type Config struct {
	HOST         string `json:"host"`
	USER         string `json:"user"`
	PASSWORD     string `json:"password"`
	DBNAME       string `json:"dbname"`
	JHUAPARTMENT string `json:"jhuapartment"`
	CRIME        string `json:"crimedata"`
	PORT         int    `json:"port"`
}

// Request ...
type Request struct {
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Crime ...
type Crime struct {
	UID         int64   `json:"id"`
	Description string  `json:"category"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// Response ...
type Response struct {
	Type     string  `json:"type"`
	AllCrime []Crime `json:"all_crime"`
	Status   int     `json:"status_code"`
}
