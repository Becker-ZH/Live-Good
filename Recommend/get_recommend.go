package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func getRecommend(dbConn Config, req Request) ([]ApartmentInfo, bool, error) {

	var res []ApartmentInfo
	var rentRange int64

	rentRange = 100

	// var number int64
	// number = 10

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return res, false, err
	}

	rent, distance, crime, convenience, saved, err := getPrefer(dbConn, req)

	// If not saved preference -> give back as defaulted list
	if !saved || rent <= 0 {
		rent = 1200.00
		distance = 0.50
		crime = 0.50
		convenience = 0.50
		saved = false
	} else { // check & give back as preferred

		if distance < 0.50 {
			distance = 0.50
		}

		if crime+convenience != 1.0 {
			crime = crime / (crime + convenience)
			convenience = convenience / (convenience)
		}
	}

	if !saved {
		res, err = getDefault(dbConn, req)
	} else {
		minRent := rent + rentRange
		maxRent := rent + rentRange

		res, err = getReduced(dbConn, req, minRent, maxRent, distance, crime, convenience)
	}

	return res, saved, err

}

func getPrefer(dbConn Config, req Request) (int64, float64, float64, float64, bool, error) {

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		return 0, 0.00, 0.00, 0.00, false, err
	}

	preference := dbConn.PREFERENCE

	var rent int64
	var distance float64
	var crime float64
	var convenience float64

	sqlStatement := `SELECT rent, distance, crime, convenience FROM ` + preference + ` WHERE username = $1;`
	// fmt.Println(sqlStatement)

	row := db.QueryRow(sqlStatement, req.UserName)

	switch err := row.Scan(
		&rent,
		&distance,
		&crime,
		&convenience,
	); err {
	case sql.ErrNoRows:
		fmt.Println(err)
		return rent, distance, crime, convenience, false, err

	case nil:
		return rent, distance, crime, convenience, true, nil

	default:
		fmt.Println(err)
		return rent, distance, crime, convenience, false, errors.New("Nothing Happended")
	}
}
