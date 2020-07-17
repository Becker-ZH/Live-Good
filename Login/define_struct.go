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
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Response ...
type Response struct {
	Status   int    `json:"status"`
	Feedback string `json:"feedback"`
}
