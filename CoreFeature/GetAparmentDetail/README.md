# Login:: Go 1.x REST API for LiveGood

**GetApartmentDetail** is a Go 1.x REST API for LiveGood.

## Features

* Get basic information for an apartment.


#### API Endpoint - POST
```URL
https://g7ndy913g7.execute-api.us-east-1.amazonaws.com/apartmentdetail/
```

#### Request Body 

```JSON
{
  "name": "Clare Court",
  "longitude": -76.6034954,
  "latitude": 39.3358207
}
```

#### Result

```JSON
Response:
{
  "name": "Clare Court",
  "longitude": -76.6034954,
  "latitude": 39.3358207,
  "apartment_info": {
    "rating": 0,
    "price": {
      "Float64": 0,
      "Valid": false
    },
    "floor_plan": {
      "String": "",
      "Valid": false
    },
    "restaurant_number": {
      "Int64": 2,
      "Valid": true
    },
    "market_number": {
      "Int64": 4,
      "Valid": true
    },
    "gas_number": {
      "Int64": 0,
      "Valid": true
    },
    "bar": {
      "Int64": 2,
      "Valid": true
    },
    "park": {
      "Int64": 3,
      "Valid": true
    },
    "gym": {
      "Int64": 0,
      "Valid": true
    },
    "hotel": {
      "Int64": 1,
      "Valid": true
    },
    "zip_code": {
      "Int64": 21218,
      "Valid": true
    }
  },
  "crime_data": null
}
```

