package main

func readAll(dbConn Config) (AllApartment, error) {

	var resp AllApartment

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return resp, err
	}

	tablename := dbConn.TABLENAME

	sqlStatement := `SELECT name, description, latitude, longitude FROM ` + tablename
	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return resp, err
	}

	var subInfo ApartmentInfo
	infos := []ApartmentInfo{}

	for rows.Next() {
		var c Coordinates

		err = rows.Scan(
			&subInfo.Name,
			&subInfo.Description,
			&c.Latitude,
			&c.Longitude,
		)
		if err != nil {
			return resp, err
		}
		subInfo.Coordinate = c
		infos = append(infos, subInfo)
	}
	err = rows.Err()
	if err != nil {
		return resp, err
	}

	resp = AllApartment{
		AllApartmentInfo: infos,
	}

	return resp, err

}
