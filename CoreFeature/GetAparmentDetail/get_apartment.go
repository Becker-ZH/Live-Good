package main

import (
	"database/sql"
)

func getApartmentInfo(dbConn Config, req Request) (ApartmentInfo, error) {

	var apartmentInfo ApartmentInfo

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return apartmentInfo, err
	}

	table := dbConn.TABLENAME
	// sqlStatement := `SELECT name, longitude, latitude, rating, price, floor_plan, restaurant, market, gas FROM ` + table
	sqlStatement := `SELECT name, longitude, latitude, address, rating, price, floorplan, restaurant, market, gas, bar, park, gym, hotel, zipcode FROM ` + table + ` ;`

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return apartmentInfo, err
	}

	var n string
	var lgt float64
	var lat float64

	for rows.Next() {
		err = rows.Scan(
			&n,
			&lgt,
			&lat,
			&apartmentInfo.Address,
			&apartmentInfo.Rating,
			&apartmentInfo.Price,
			&apartmentInfo.FloorPlan,
			&apartmentInfo.Restaurant,
			&apartmentInfo.Market,
			&apartmentInfo.Gas,
			&apartmentInfo.Bar,
			&apartmentInfo.Park,
			&apartmentInfo.Gym,
			&apartmentInfo.Hotel,
			&apartmentInfo.Zipcode,
		)
		if err != nil {
			return apartmentInfo, err
		}

		if n == req.Name || (lgt == req.Longitude && lat == req.Latitude) {

			dbConnCrime, err := getCrimeConfigDetailes()
			if err != nil {
				return apartmentInfo, err
			}

			c, err := getCrime(dbConnCrime, req)
			if err != nil {
				apartmentInfo.CrimeNum = Fake{
					Int64: 0,
					Valid: false,
				}
			} else {
				apartmentInfo.CrimeNum = c
			}

			if apartmentInfo.Restaurant.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Restaurant.Int64
			}
			if apartmentInfo.Market.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Market.Int64
			}
			if apartmentInfo.Gas.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Gas.Int64
			}
			if apartmentInfo.Bar.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Bar.Int64
			}
			if apartmentInfo.Gym.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Gym.Int64
			}
			if apartmentInfo.Park.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Park.Int64
			}
			if apartmentInfo.Hotel.Valid {
				apartmentInfo.UtilityNum.Int64 += apartmentInfo.Hotel.Int64
			}

			if apartmentInfo.UtilityNum.Int64 > 0 {
				apartmentInfo.UtilityNum.Valid = true
			} else {
				apartmentInfo.UtilityNum.Valid = false
			}

			return apartmentInfo, nil
		}
	}

	err = rows.Err()
	if err != nil {
		return apartmentInfo, err
	}

	empty := ApartmentInfo{
		Rating:    0,
		Price:     sql.NullFloat64{},
		FloorPlan: sql.NullString{},
	}

	return empty, nil
}
