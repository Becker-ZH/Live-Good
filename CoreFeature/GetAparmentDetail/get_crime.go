package main

func getCrime(dbConn Config, req Request) (Fake, error) {
	var count int64
	var res Fake

	miles := 0.5

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return res, err
	}

	table := dbConn.TABLENAME
	sqlStatement := `SELECT id, latitude, longitude FROM ` + table

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	var uid int64
	var lat float64
	var lon float64

	for rows.Next() {
		err = rows.Scan(
			&uid,
			&lat,
			&lon,
		)
		if err != nil {
			return res, err
		}

		if earthDistance(req.Latitude, req.Longitude, lat, lon) <= miles {
			count++
		}
	}

	err = rows.Err()
	if err != nil {
		return res, err
	}

	res = Fake{
		Int64: count,
		Valid: true,
	}

	return res, nil
}
