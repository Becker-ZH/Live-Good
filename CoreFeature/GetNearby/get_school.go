package main

import (
	"math"
)


//naive distance algorithm.
func earthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137 //define radius of earth
	rad := math.Pi / 180.0 //define radian
	//calculate distance
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius * 0.621371
}

func getSchool(dbConn Config, req Request) (School, error) {

	miles := 10.0

	var resp School

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return resp, err
	}

	table := dbConn.TABLENAME
	sqlStatement := `SELECT name, longitude, latitude FROM ` + table

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return resp, err
	}

	findMatch := false
	miniDist := 1.7976931348623157e+308
	var miniResp School

	for rows.Next() {
		err = rows.Scan(
			&resp.Name,
			&resp.Longitude,
			&resp.Latitude,
		)
		if err != nil {
			return resp, err
		}

		distance := earthDistance(req.Latitude, req.Longitude, resp.Latitude, resp.Longitude)
		if distance <= miles {
			if findMatch == false {
				findMatch = true
				miniDist = distance
				miniResp = resp
			} else {
				if distance < miniDist {
					miniDist = distance
					miniResp = resp
				}
			}
		}
	}

	err = rows.Err()
	if err != nil {
		return resp, err
	}

	if findMatch == false {
		miniResp.Name = ""
	}

	return miniResp, nil

}

func getSchoolList(dbConn Config) ([]School, error) {
	var schoolList []School

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return schoolList, err
	}

	table := dbConn.TABLENAME
	sqlStatement := `SELECT name, longitude, latitude FROM ` + table

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return schoolList, err
	}

	var tmp School

	for rows.Next() {
		err = rows.Scan(
			&tmp.Name,
			&tmp.Longitude,
			&tmp.Latitude,
		)
		if err != nil {
			return schoolList, err
		}

		schoolList = append(schoolList, tmp)
	}

	err = rows.Err()
	if err != nil {
		return schoolList, err
	}

	return schoolList, nil
}
