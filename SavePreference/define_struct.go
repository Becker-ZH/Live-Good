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
	Username    string  `json:"username"`
	Rent        int64   `json:"rent"`
	Bedroom     int64   `json:"bedroom"`
	Bathroom    int64   `json:"bathroom"`
	Distance    float64 `json:"distance"`
	Crime       float64 `json:"crime"`
	Convenience float64 `json:"convenience"`
}

// Response ...
type Response struct {
	Status   int    `json:"status"`
	Feedback string `json:"feedback"`
}
