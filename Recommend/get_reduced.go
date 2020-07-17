package main

import "sort"

func getReduced(dbConn Config, req Request, minRent int64, maxRent int64, distance float64, crime float64, convenience float64) ([]ApartmentInfo, error) {
	var res Apartments

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return res, err
	}

	tablename := dbConn.JHUAPARTMENT

	sqlStatement := `SELECT index, name, latitude, longitude, address, rating, zipcode, price, floorplan, restaurant, market, gas, bar, park, gym, hotel FROM ` + tablename

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	var subInfo ApartmentInfo
	var restaurant int64
	var market int64
	var gas int64
	var bar int64
	var park int64
	var gym int64
	var hotel int64

	for rows.Next() {

		err = rows.Scan(
			&subInfo.ID,
			&subInfo.Name,
			&subInfo.Latitude,
			&subInfo.Longitude,
			&subInfo.Address,
			&subInfo.Rating,
			&subInfo.Zipcode,
			&subInfo.Price,
			&subInfo.FloorPlan,
			&restaurant,
			&market,
			&gas,
			&bar,
			&park,
			&gym,
			&hotel,
		)
		if err != nil {
			return res, err
		}

		if subInfo.Price.Valid && subInfo.Price.Float64 < float64(maxRent) {
			if earthDistance(subInfo.Latitude, subInfo.Longitude, req.Latitude, req.Longitude) <= distance {

				var numCrime int64
				numCrime, err = getCrimeNum(dbConn, subInfo.Latitude, subInfo.Longitude, distance)
				if err != nil {
					return res, err
				}

				con := float64(restaurant + market + gas + bar + park + gym + hotel)
				lifeIndex := con*convenience - crime*float64(numCrime)*0.01
				subInfo.LifeIndex = lifeIndex

				subInfo.UntilityNum.Int64 = restaurant + market + gas + bar + park + gym + hotel
				if subInfo.UntilityNum.Int64 == 0 {
					subInfo.UntilityNum.Valid = false
				} else {
					subInfo.UntilityNum.Valid = true
				}

				subInfo.CrimNum.Int64 = numCrime
				if subInfo.CrimNum.Int64 > 0 {
					subInfo.CrimNum.Valid = true
				} else {
					subInfo.CrimNum.Valid = false
				}

				res = append(res, subInfo)
			}
		}
	}

	err = rows.Err()
	if err != nil {
		return res, err
	}

	sort.Sort(res)

	var rCut []ApartmentInfo
	count := 0

	for _, r := range res {
		if count < 8 {
			rCut = append(rCut, r)
		} else {
			break
		}
	}

	return rCut, err

}

func getCrimeNum(dbConn Config, lati float64, longi float64, distance float64) (int64, error) {

	var c int64

	c = 0.00

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return c, err
	}

	tablename := dbConn.CRIME

	sqlStatement := `SELECT id, latitude, longitude FROM ` + tablename

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return c, err
	}

	var uid int64
	var latitude float64
	var longitude float64

	for rows.Next() {

		err = rows.Scan(
			&uid,
			&latitude,
			&longitude,
		)
		if err != nil {
			return c, err
		}

		if earthDistance(lati, longi, latitude, longitude) <= distance {
			c++
		}
	}
	err = rows.Err()
	if err != nil {
		return c, err
	}

	return c, err

}
