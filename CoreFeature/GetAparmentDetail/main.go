package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, req Request) (Response, error) {
	var response Response

	// Connect to table apartment_detail
	dbConn, err := getApartmentConfigDetailes()
	if err != nil {
		return response, err
	}

	response.Name = req.Name
	response.Longitude = req.Longitude
	response.Latitude = req.Latitude

	response.Apartment, err = getApartmentInfo(dbConn, req)
	if err != nil {
		return response, err
	}

	// Connect to crime_data
	// dbConn, err = getCrimeConfigDetailes()
	// if err != nil {
	// 	return response, err
	// }

	// response.CrimeData, err = getCrimeDataList(dbConn, req)
	// if err != nil {
	// 	return response, err
	// }

	//connect to shop_data
	//dbConn, err = getShopConfigDetailes()
	//if err != nil{
	//    return response, err
	//}

	//response.ShopData, err = getShopDataList(dbConn, req)
	//if err != nil{
	//    return response, err
	//}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
