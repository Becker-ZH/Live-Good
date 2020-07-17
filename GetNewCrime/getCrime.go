package main

import (
	"database/sql"
	"errors"
)

func getReqID(dbConn Config, req Request) (int64, error) {
	var id int64

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return id, err
	}

	tablename := dbConn.JHUAPARTMENT

	latitude := req.Latitude
	longitude := req.Longitude

	sqlStatement := `SELECT index FROM ` + tablename + ` WHERE latitude = $1 AND longitude = $2;`
	// fmt.Println(sqlStatement, latitude, longitude)

	row := db.QueryRow(sqlStatement, latitude, longitude)

	switch err := row.Scan(
		&id,
	); err {
	case sql.ErrNoRows:
		return 0, err

	case nil:
		return id, nil

	default:
		return 0, errors.New("Nothing Happended")
	}
}

func getCrimeList(dbConn Config, req Request) ([]Crime, error) {
	var c []Crime

	miles := 0.45
	count := 0

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return c, err
	}

	_, err = getReqID(dbConn, req)
	if err != nil {
		return c, err
	}

	tablename := dbConn.CRIME

	sqlStatement := `SELECT id, description, latitude, longitude FROM ` + tablename

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return c, err
	}

	var subInfo Crime

	for rows.Next() {

		err = rows.Scan(
			&subInfo.UID,
			&subInfo.Description,
			&subInfo.Latitude,
			&subInfo.Longitude,
		)
		if err != nil {
			return c, err
		}

		if count <= 20 {
			if earthDistance(subInfo.Latitude, subInfo.Longitude, req.Latitude, req.Longitude) <= miles {
				c = append(c, subInfo)
				count = count + 1
			}
		}
	}
	err = rows.Err()
	if err != nil {
		return c, err
	}

	return c, err
}
