package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, req Request) (Response, error) {

	var response Response

	// fmt.Println(request)

	// requestLatitude, err := strconv.ParseFloat(request.QueryStringParameters["latitude"], 64)
	// if err != nil {
	// 	return response, err
	// }

	// requestLongitude, err := strconv.ParseFloat(request.QueryStringParameters["longitude"], 64)
	// if err != nil {
	// 	return response, err
	// }

	// requestType := request.QueryStringParameters["type"]

	// req := Request{
	// 	Type:      requestType,
	// 	Longitude: requestLongitude,
	// 	Latitude:  requestLatitude,
	// }

	dbConn, err := getSchoolConfigDetailes()
	if err != nil {
		return response, err
	}

	schoolInfo, err := getSchool(dbConn, req)
	if err != nil {
		return response, err
	}

	if len(schoolInfo.Name) > 0 {
		response.SchoolInfo = schoolInfo
		response.Type = req.Type
		response.SchoolList = nil
		response.FindMatch = true
	} else {
		response.FindMatch = false
		response.Type = req.Type

		response.SchoolList, err = getSchoolList(dbConn)
		if err != nil {
			return response, err
		}

		return response, nil
	}

	dbConn, err = getApartmentConfigDetailes()
	if err != nil {
		return response, err
	}

	apartments, err := readAll(dbConn, req)
	if err != nil {
		return response, err
	}

	response.Apartments = apartments

	return response, err

}

func main() {
	lambda.Start(handler)
}
