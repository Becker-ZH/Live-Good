package main

import "github.com/aws/aws-lambda-go/lambda"

func handler(req Request) (Response, error) {

	var response Response

	dbConn, err := getConfigDetailes()
	if err != nil {
		response = Response{
			Status:   404,
			Feedback: "Error with DB connection",
		}
		return response, err
	}

	response, err = save(dbConn, req)
	return response, err

}

func main() {
	lambda.Start(handler)
}
