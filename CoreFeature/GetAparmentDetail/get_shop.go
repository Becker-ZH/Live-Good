package main

// func getShopDataList(dbConn Config, req Request) ([]Shop, error) {
// 	var shopDataList []Shop
// 	miles := 0.1

// 	db, err := connectDB(dbConn)
// 	defer db.Close()

// 	if err != nil {
// 		return shopDataList, err
// 	}

// 	table := dbConn.TABLENAME
// 	sqlStatement := `SELECT neighborhood, longitude, latitude, location FROM ` + table

// 	rows, err := db.Query(sqlStatement)
// 	defer rows.Close()

// 	if err != nil {
// 		return shopDataList, err
// 	}

// 	var shopRow Shop
// 	for rows.Next() {
// 		err = rows.Scan(
// 			&shopRow.Name,
// 			&shopRow.Type,
// 			&shopRow.Latitude,
// 			&shopRow.Location,
// 		)
// 		if err != nil {
// 			return shopDataList, err
// 		}

// 		if earthDistance(req.Latitude, req.Longitude, shopRow.Latitude, shopRow.Longitude) <= miles {
// 			shopDataList = append(shopDataList, shopRow)
// 		}
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return shopDataList, err
// 	}

// 	return shopDataList, nil
// }
