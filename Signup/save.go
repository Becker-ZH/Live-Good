package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func save(dbConn Config, req Request) (Response, error) {
	var response Response

	first := req.FirstName
	last := req.LastName

	user := req.UserName
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)

	if err != nil {
		response = Response{
			Status:   404,
			Feedback: "Error with bcrypt",
		}
		return response, err
	}

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
	sqlStatement := `SELECT username FROM ` + tablename + ` WHERE username=$1;`

	row := db.QueryRow(sqlStatement, user)

	var current string
	switch err := row.Scan(&current); err {
	case sql.ErrNoRows:
		sqlStatement = `INSERT INTO ` + tablename + ` VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
		_, err = db.Exec(
			sqlStatement,
			user,
			hashed,
			first,
			last,
			0,
			2,
			2,
			0.5,
			0.5,
			0.5,
		)

		if err != nil {
			response = Response{
				Status:   404,
				Feedback: "Error with insert",
			}
			return response, err
		}

		response = Response{
			Status:   200,
			Feedback: "Sign up success",
		}
		return response, err

	case nil:
		response = Response{
			Status:   404,
			Feedback: "User name already exists",
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
