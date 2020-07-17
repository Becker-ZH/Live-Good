package main

import "database/sql"

func getDefault(dbConn Config, req Request) ([]ApartmentInfo, error) {
	var res []ApartmentInfo

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return res, err
	}

	miles := 0.5
	// count := 0

	tablename := dbConn.JHUAPARTMENT

	sqlStatement := `SELECT index, name, latitude, longitude, address, rating, zipcode, price, floorplan, restaurant, market, gas, bar, park, gym, hotel FROM ` + tablename

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	var subInfo ApartmentInfo
	var restaurant sql.NullInt64
	var market sql.NullInt64
	var gas sql.NullInt64
	var bar sql.NullInt64
	var park sql.NullInt64
	var gym sql.NullInt64
	var hotel sql.NullInt64

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

		if earthDistance(subInfo.Latitude, subInfo.Longitude, req.Latitude, req.Longitude) <= miles {

			if restaurant.Valid {
				subInfo.UntilityNum.Int64 += restaurant.Int64
			}
			if market.Valid {
				subInfo.UntilityNum.Int64 += market.Int64
			}
			if gas.Valid {
				subInfo.UntilityNum.Int64 += gas.Int64
			}
			if bar.Valid {
				subInfo.UntilityNum.Int64 += bar.Int64
			}
			if gym.Valid {
				subInfo.UntilityNum.Int64 += gym.Int64
			}
			if park.Valid {
				subInfo.UntilityNum.Int64 += park.Int64
			}
			if hotel.Valid {
				subInfo.UntilityNum.Int64 += hotel.Int64
			}

			if subInfo.UntilityNum.Int64 > 0 {
				subInfo.UntilityNum.Valid = true
			} else {
				subInfo.UntilityNum.Valid = false
			}

			var numCrime int64
			numCrime, err = getCrimeNum(dbConn, subInfo.Latitude, subInfo.Longitude, miles)
			if err != nil {
				return res, err
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

	err = rows.Err()
	if err != nil {
		return res, err
	}

	return res, err

}
