package main

import "database/sql"

func update(dbConn Config, req Request) (Response, error) {
	var response Response

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		response = Response{
			Status:   404,
			Feedback: "Error with DB connection",
		}
		return response, err
	}

	tablename := dbConn.TABLENAME
	sqlStatement := `SELECT username FROM ` + tablename + ` WHERE username=$1;`
	row := db.QueryRow(sqlStatement, req.Username)

	var u string
	switch err := row.Scan(&u); err {

	case sql.ErrNoRows:
		response = Response{
			Status:   404,
			Feedback: "User not exsits",
		}
		return response, err

	case nil:
		sqlStatement = `UPDATE ` + tablename + ` SET rent=$1, bedroom=$2, bathroom=$3, distance=$4, crime=$5, convenience=$6 WHERE username=$7;`
		_, err = db.Exec(
			sqlStatement,
			req.Rent,
			req.Bedroom,
			req.Bathroom,
			req.Distance,
			req.Crime,
			req.Convenience,
			req.Username,
		)
		if err != nil {
			response = Response{
				Status:   404,
				Feedback: "Error with update DB",
			}
			return response, err
		}

		response = Response{
			Status:   200,
			Feedback: "Update preference",
		}
		return response, err

	default:
		response = Response{
			Status:   404,
			Feedback: "Nothing happened",
		}
		return response, err
	}
}
