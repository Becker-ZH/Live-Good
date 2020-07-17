# CoreFeature :: Go 1.x REST API for LiveGood

**CoreFeature** is a Go 1.x REST API for LiveGood.

## Features - 1

* Give a user's coordinate data and return his nearest school information and nearby aparmtents information.


#### API Endpoint
POST:
```URL
https://9hk8gkmxfe.execute-api.us-east-1.amazonaws.com/nearby/
```

GET:
```URL
https://g486jas5j8.execute-api.us-east-1.amazonaws.com/nearby/
```

#### Request Body - Raw
* `type`: (string, optional)
* `longitude`: (float64)
* `latitude`: (float64)

```JSON
{
  "type": "nearby",
  "longitude": -76.6205235,
  "latitude": 39.3289957
}
```

#### Result

- `type`: denotes the request type
- `find_match`: if no school is found, this will be `false`; otherwise, `true`
- `nearby_school`: nearest school information, e.g. 10 miles
  - `name`: school name
  - `latitude`: (double, float64)
  - `longitude`: (double, float64) 
- `nearby_aparmtents`: nearby apartments within a distanct, e.g. 5 miles
  - `name`: school name
  - `latitude`: (double, float64)
  - `longitude`: (double, float64) 
  - `rating`: rate of the apartment
- `school_list`: if `find_match` is true, it will give back `null`; otherwise, it will give back a whole list of schools

##### Response:
- `find_match` == True

```JSON

{
  "type": "nearby",
  "find_match": true,
  "nearby_school": {
    "name": "Johns Hopkins University",
    "longitude": -76.6205235,
    "latitude": 39.3289957
  },
  "nearby_apartments": [
    {
      "name": "The Marylander Apartment Homes",
      "longitude": -76.6157407,
      "latitude": 39.330454,
      "rating": 2.9
    },
    {
      "name": "Guilford Manor Apartments",
      "longitude": -76.61875169999999,
      "latitude": 39.3336055,
      "rating": 4.5
    },
    {
      "name": "The Academy on Charles",
      "longitude": -76.61837729999999,
      "latitude": 39.334015,
      "rating": 2.8
    },
    ...
  ],
  "school_list": null
}
```
- `find_match` == False

```JSON
{
  "type": "nearby",
  "find_match": false,
  "nearby_school": {
    "name": "",
    "longitude": 0,
    "latitude": 0
  },
  "nearby_apartments": null,
  "school_list": [
    {
      "name": "Johns Hopkins University",
      "longitude": -76.6205235,
      "latitude": 39.3289957
    }
  ]
}
```


## Features - 2
- This API will give back the detail information of the aparmtent, e.g. it's prices, floor plans, ratings, nearby crime data, and restaurants, markets, gas stations.

#### API Endpoint
POST:
```URL
https://g7ndy913g7.execute-api.us-east-1.amazonaws.com/apartmentdetail/
```

#### Request Body - Raw
- name: (string, required) the name of the apartment
- longitude: (float64, required) 
- latitude: (float64, required)

```JSON
{
  "name": "The Academy on Charles",
  "longitude": -76.6183773,
  "latitude": 39.334015
}
```

#### Result
- name: (string)
- longitude: (float4)
- latitude: (float64)
- apartment_info
  - rating: (float64)
  - price: (sql.NullString)
  - floor_plan: (sql.NullSring)
  - restaurant_number, market_number, gas_number: (sql.NullInt)
- crime_data:
  - neighborhood: (string)
  - longitude, latitude: (float64)
  - location: (string)

```JSON
{
  "name": "The Academy on Charles",
  "longitude": -76.6183773,
  "latitude": 39.334015,
  "apartment_info": {
    "rating": 2.8,
    "price": {
      "String": "1165",
      "Valid": true
    },
    "floor_plan": {
      "String": "1b1b;3b2b;4b4b",
      "Valid": true
    },
    "restaurant_number": {
      "Int64": 13,
      "Valid": true
    },
    "market_number": {
      "Int64": 7,
      "Valid": true
    },
    "gas_number": {
      "Int64": 0,
      "Valid": true
    }
  },
  "crime_data": [
    {
      "neighborhood": "GUILFORD",
      "longitude": -76.61750211020001,
      "latitude": 39.3348278017,
      "location": "0 SAINT MARTINS RD"
    },
    {
      "neighborhood": "TUSCANY-CANTERBURY",
      "longitude": -76.618962,
      "latitude": 39.333269,
      "location": "0 W UNIVERSITY PKWY"
    },
    {
      "neighborhood": "GUILFORD",
      "longitude": -76.617863,
      "latitude": 39.334745,
      "location": "0 SAINT MARTINS RD"
    },
    ...
  ]
}

```

- More test data

```JSON
{
  "name": "Hopkins House Apartments",
  "longitude": -76.6230081,
  "latitude": 39.3365522
}
{
  "name": "The Marylander Apartment Homes",
  "longitude": -76.6157407,
  "latitude": 39.330454
}
{
  "name": "Nine East 33rd",
  "longitude": -76.6165309,
  "latitude": 39.3278957
}
{
  "name": "The Social North Charles",
  "longitude": -76.6185808,
  "latitude": 39.3362825
}
```