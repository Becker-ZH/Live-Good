package main

import "github.com/aws/aws-lambda-go/lambda"

func handler(req Request) (Response, error) {

	var response Response

	dbConn, err := getConfigDetailes()
	if err != nil {
		response = Response{
			Status: 404,
		}
		return response, err
	}

	response, err = read(dbConn, req)
	if err != nil {
		response = Response{
			Status: 404,
		}
		return response, err
	}

	return response, err

}

func main() {
	lambda.Start(handler)
}
