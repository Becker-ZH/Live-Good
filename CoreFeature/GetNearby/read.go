package main

import "fmt"

func readAll(dbConn Config, req Request) ([]ApartmentInfo, error) {
	miles := 5.00

	var resp []ApartmentInfo

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return resp, err
	}

	tablename := dbConn.TABLENAME

	sqlStatement := `SELECT name, latitude, longitude, rating FROM ` + tablename
	fmt.Println(sqlStatement)
	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return resp, err
	}

	var subInfo ApartmentInfo
	infos := []ApartmentInfo{}

	for rows.Next() {

		err = rows.Scan(
			&subInfo.Name,
			&subInfo.Latitude,
			&subInfo.Longitude,
			&subInfo.Rate,
		)
		if err != nil {
			return resp, err
		}

		if earthDistance(subInfo.Latitude, subInfo.Longitude, req.Latitude, req.Longitude) <= miles {
			infos = append(infos, subInfo)
		}

	}
	err = rows.Err()
	if err != nil {
		return resp, err
	}

	resp = infos

	return resp, err

}
