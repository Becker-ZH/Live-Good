# FirstFeature :: Go 1.x REST API for LiveGood

**FirstFeature** is a Go 1.x REST API for LiveGood.

## Features

* Read all apartment data from table: `apartment_info` in DB: `LiveGood`

### <ReadAll\>
The `readAll` method receives am empty request body from the front end and sends a `POST` request to `apartment_info` table and return result as a `JSON` object containing all the apartment information in the table.

#### API Endpoint
```URL
https://deqjn0w7h0.execute-api.us-east-1.amazonaws.com/dev/
```

#### Request Body 

```JSON
{
  
}
```

#### Result

- `all_apartment`: denotes the request type
  - `id`: (string) the unique of ID of apartment
  - `title`: (string) the title/name of the apartment
  - `decs`: (string) the description of the apartment
  - `coordinates`: the coordinates struct for the apartment
    - `latitude`: (double, float64)
    - `longitude`: (double, float64) 

```JSON
Response:
{
  "all_apartment": [
    {
      "id": "",
      "title": "Academy on Charles",
      "desc": "",
      "coordinates": {
        "latitude": 39.334,
        "longitude": -76.6184
      }
    },
    {
      "id": "",
      "title": "Hopkins House",
      "desc": "",
      "coordinates": {
        "latitude": 39.3366,
        "longitude": -76.623
      }
    }
  ]
}
```

