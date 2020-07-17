package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func read(dbConn Config, req Request) (Response, error) {
	var response Response

	db, err := connectDB(dbConn)
	defer db.Close()

	if err != nil {
		response = Response{
			Status:   404,
			Feedback: "Error with connecting DB",
		}
		return response, err
	}

	tablename := dbConn.TABLENAME
	sqlStatement := `SELECT username, password FROM ` + tablename + ` WHERE username=$1;`

	row := db.QueryRow(sqlStatement, req.UserName)
	pwdByte := []byte(req.Password)

	var curUser string
	var curPwd string

	switch err := row.Scan(
		&curUser,
		&curPwd,
	); err {
	case sql.ErrNoRows:
		response = Response{
			Status:   404,
			Feedback: "User not exists",
		}
		return response, err

	case nil:
		err = bcrypt.CompareHashAndPassword([]byte(curPwd), pwdByte)

		if err != nil {
			response = Response{
				Status:   404,
				Feedback: "Login fail - password error",
			}
			return response, err
		}
		response = Response{
			Status:   200,
			Feedback: "Login success",
		}
		return response, err

	default:
		response = Response{
			Status:   404,
			Feedback: "Nothing happens",
		}
		return response, err
	}
}
