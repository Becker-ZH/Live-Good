package main

import (
	"database/sql"
	"errors"
)

func read(dbConn Config, req Request) (Response, error) {

	var resp Response

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return resp, err
	}

	var rent int64
	var bedroom int64
	var bathroom int64
	var distance float64
	var crime float64
	var convenience float64

	tablename := dbConn.TABLENAME

	sqlStatement := `SELECT rent, bedroom, bathroom, distance, crime, convenience FROM ` + tablename + ` WHERE username = $1;`

	row := db.QueryRow(sqlStatement, req.Username)

	switch err := row.Scan(
		&rent,
		&bedroom,
		&bathroom,
		&distance,
		&crime,
		&convenience,
	); err {
	case sql.ErrNoRows:
		return resp, err

	case nil:
		resp = Response{
			Status:      200,
			Rent:        rent,
			Bedroom:     bedroom,
			Bathroom:    bathroom,
			Distance:    distance,
			Crime:       crime,
			Convenience: convenience,
		}
		return resp, nil

	default:
		return resp, errors.New("Nothing Happended")
	}

}
