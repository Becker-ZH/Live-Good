# Recommend:: Go 1.x REST API for LiveGood

**Recommend** is a Go 1.x REST API for LiveGood.

## Features

* Reommendation System.


#### API Endpoint - POST
```URL
https://3ktr1y4ry5.execute-api.us-east-1.amazonaws.com/recommend/
```

#### Request Body 

```JSON
{
  "username": "9088",
  "latitude": 39.3289957,
  "longitude": -76.6205235
}
```

#### Result

```JSON
Response:
{
  "type": "get_recommend",
  "saved_prefernce": true,
  "recommend": [
    {
      "id": 12,
      "name": "Jefferson House Apartments",
      "latitude": 39.327062,
      "longitude": -76.616472,
      "address": "4 E 32nd St, Baltimore",
      "rating": 4.6,
      "price": {
        "Float64": 880,
        "Valid": true
      },
      "floor_plan": {
        "String": "",
        "Valid": false
      },
      "zip_code": {
        "Int64": 21218,
        "Valid": true
      },
      "life_index": -739
    },
    {
      "id": 5,
      "name": "The Marylander Apartment Homes",
      "latitude": 39.330454,
      "longitude": -76.6157407,
      "address": "3501 St Paul St, Baltimore",
      "rating": 2.8,
      "price": {
        "Float64": 1188.84,
        "Valid": true
      },
      "floor_plan": {
        "String": "",
        "Valid": false
      },
      "zip_code": {
        "Int64": 21218,
        "Valid": true
      },
      "life_index": -741
    },
    {
      "id": 18,
      "name": "Charles & Blackstone Apartments",
      "latitude": 39.3286441,
      "longitude": -76.6171646,
      "address": "3333 N Charles St, Baltimore",
      "rating": 2.8,
      "price": {
        "Float64": 1250,
        "Valid": true
      },
      "floor_plan": {
        "String": "",
        "Valid": false
      },
      "zip_code": {
        "Int64": 21218,
        "Valid": true
      },
      "life_index": -741
    },
    {
      "id": 10,
      "name": "The Academy on Charles",
      "latitude": 39.3340036,
      "longitude": -76.6189778,
      "address": "3700 N Charles St, Baltimore",
      "rating": 2.9,
      "price": {
        "Float64": 1300,
        "Valid": true
      },
      "floor_plan": {
        "String": "",
        "Valid": false
      },
      "zip_code": {
        "Int64": 21218,
        "Valid": true
      },
      "life_index": -764
    },
    {
      "id": 11,
      "name": "Ambassador Apartments LLC",
      "latitude": 39.3355576,
      "longitude": -76.61992099999999,
      "address": "3811 Canterbury Rd, Baltimore",
      "rating": 3,
      "price": {
        "Float64": 1507.875,
        "Valid": true
      },
      "floor_plan": {
        "String": "",
        "Valid": false
      },
      "zip_code": {
        "Int64": 21218,
        "Valid": true
      },
      "life_index": -769
    }
  ],
  "status_code": 200
}
```

