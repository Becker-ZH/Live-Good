package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, req Request) (Response, error) {

	var response Response

	dbConn, err := getApartmentConfigDetailes()
	if err != nil {
		return response, err
	}

	var near []ApartmentInfo
	near, saved, err := getRecommend(dbConn, req)

	if err != nil {
		response = Response{
			Type:          "get_recommend",
			SavePrefernce: saved,
			Recommend:     near,
			Status:        404,
		}
		return response, err
	}

	response = Response{
		Type:          "get_recommend",
		SavePrefernce: saved,
		Recommend:     near,
		Status:        200,
	}

	return response, err

}

func main() {
	lambda.Start(handler)
}
