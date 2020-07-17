package main

// Config ... Configuration parameters for database connection
type Config struct {
	HOST         string `json:"host"`
	USER         string `json:"user"`
	PASSWORD     string `json:"password"`
	DBNAME       string `json:"dbname"`
	JHUAPARTMENT string `json:"jhuapartment"`
	JHUPLACE     string `json:"jhuplace"`
	APTPLACE     string `json:"aptplace"`
	PORT         int    `json:"port"`
}

// Request ...
type Request struct {
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Nearby ...
type Nearby struct {
	UID       int64   `json:"id"`
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Response ...
type Response struct {
	Type      string   `json:"type"`
	AllNearby []Nearby `json:"all_nearby"`
	Status    int      `json:"status_code"`
}
