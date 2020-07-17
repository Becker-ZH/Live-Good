package main

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
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// School ...
type School struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// ApartmentInfo ... Apartment information
type ApartmentInfo struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Rate      float64 `json:"rating"`
}

// Response ...
type Response struct {
	Type       string          `json:"type"`
	FindMatch  bool            `json:"find_match"`
	SchoolInfo School          `json:"nearby_school"`
	Apartments []ApartmentInfo `json:"nearby_apartments"`
	SchoolList []School        `json:"school_list"`
}

// APIGatewayProxyRequest ...
type APIGatewayProxyRequest struct {
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
