package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func getReqID(dbConn Config, req Request) (int64, string, error) {
	var id int64
	var name string

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return 0, name, err
	}

	tablename := dbConn.JHUAPARTMENT

	latitude := req.Latitude
	longitude := req.Longitude

	sqlStatement := `SELECT index, name FROM ` + tablename + ` WHERE latitude = $1 AND longitude = $2;`

	row := db.QueryRow(sqlStatement, latitude, longitude)

	switch err := row.Scan(
		&id,
		&name,
	); err {
	case sql.ErrNoRows:
		return 0, name, err

	case nil:
		fmt.Println(name)
		return id, name, nil

	default:
		return 0, name, errors.New("Nothing Happended")
	}
}

func getCombinedList(dbConn Config, req Request) ([]Nearby, error) {
	var n []Nearby

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return n, err
	}

	jhuApartment := dbConn.JHUAPARTMENT
	jhuPlace := dbConn.JHUPLACE
	aptPlace := dbConn.APTPLACE

	_, reqName, err := getReqID(dbConn, req)
	if err != nil {
		return n, err
	}

	sqlStatement := `SELECT PL.index, PL.name, PL.type, PL.address, PL.latitude, PL.longitude FROM ` + jhuPlace + ` AS PL, ` + jhuApartment + ` AS AP, ` + aptPlace + ` AS AL WHERE AL.aptname = AP.name AND AL.placename = PL.name and AP.name = $1;`

	rows, err := db.Query(sqlStatement, reqName)
	defer rows.Close()

	if err != nil {
		return n, err
	}

	var subN Nearby

	for rows.Next() {

		err = rows.Scan(
			&subN.UID,
			&subN.Name,
			&subN.Category,
			&subN.Address,
			&subN.Latitude,
			&subN.Longitude,
		)
		if err != nil {
			return n, err
		}

		n = append(n, subN)

	}
	err = rows.Err()
	if err != nil {
		return n, err
	}

	return n, nil
}
