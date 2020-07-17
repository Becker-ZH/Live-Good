# updateCrime:: Go 1.x REST API for LiveGood

**updateCrime** is a Go 1.x REST API for LiveGood.

## Features

* Using `jhuapartment` and `crime_data` database tables to fetch nearby crime incidents records within the past thirteen years.

### <getReqID\>
The `getReqID` method read `latitude` and `longtitude` given from the request and find the corresponding apartment ID in `jhuapartment`.

### <getCrime\>
The `getCrime` method search nearby crime incidents from `crime_data` table.

#### API Endpoint - POST
```URL
https://sxetiwcj76.execute-api.us-east-1.amazonaws.com/aptCrime/
```

#### Request Body 

```JSON
{
  "type": "crime",
  "latitude": 39.3357969,
  "longtitude": -76.6035516
}
```
- `type` is negligible.

#### Result

```JSON
{
  "type": "get_crime",
  "all_crime": [
    {
      "id": 1663,
      "category": "LARCENY",
      "latitude": 39.333702,
      "longtitude": -76.609951
    },
    {
      "id": 12287,
      "category": "ROBBERY - CARJACKING",
      "latitude": 39.336794890300006,
      "longtitude": -76.6103462577
    },
    {
      "id": 24271,
      "category": "LARCENY",
      "latitude": 39.3347214222,
      "longtitude": -76.61061881020001
    },
    {
      "id": 29756,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.3336548758,
      "longtitude": -76.6099320358
    },
    {
      "id": 39034,
      "category": "BURGLARY",
      "latitude": 39.333647,
      "longtitude": -76.609951
    },
    {
      "id": 54366,
      "category": "BURGLARY",
      "latitude": 39.334780387100004,
      "longtitude": -76.6106191539
    },
    {
      "id": 63997,
      "category": "ROBBERY - CARJACKING",
      "latitude": 39.3334746718,
      "longtitude": -76.61030052310001
    },
    {
      "id": 65286,
      "category": "BURGLARY",
      "latitude": 39.3345737613,
      "longtitude": -76.6106886572
    },
    {
      "id": 85146,
      "category": "BURGLARY",
      "latitude": 39.3360185698,
      "longtitude": -76.6106983653
    },
    {
      "id": 85147,
      "category": "BURGLARY",
      "latitude": 39.3360185698,
      "longtitude": -76.6106983653
    },
    {
      "id": 85352,
      "category": "BURGLARY",
      "latitude": 39.3373637707,
      "longtitude": -76.6106714693
    },
    {
      "id": 86649,
      "category": "BURGLARY",
      "latitude": 39.333318945500004,
      "longtitude": -76.6100800145
    },
    {
      "id": 95668,
      "category": "AUTO THEFT",
      "latitude": 39.333751284499996,
      "longtitude": -76.6101788284
    },
    {
      "id": 96306,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.336929999999995,
      "longtitude": -76.61064300000001
    },
    {
      "id": 105145,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.333647,
      "longtitude": -76.609993
    },
    {
      "id": 106446,
      "category": "BURGLARY",
      "latitude": 39.336929999999995,
      "longtitude": -76.61064300000001
    },
    {
      "id": 106696,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.336929999999995,
      "longtitude": -76.61064300000001
    },
    {
      "id": 128098,
      "category": "LARCENY",
      "latitude": 39.3355368546,
      "longtitude": -76.6106980694
    },
    {
      "id": 139637,
      "category": "BURGLARY",
      "latitude": 39.3336546943,
      "longtitude": -76.6099571124
    },
    {
      "id": 145037,
      "category": "BURGLARY",
      "latitude": 39.334653,
      "longtitude": -76.61061600000001
    },
    {
      "id": 145206,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.334653,
      "longtitude": -76.61061600000001
    },
    {
      "id": 152197,
      "category": "ROBBERY - CARJACKING",
      "latitude": 39.33635,
      "longtitude": -76.61063399999999
    },
    {
      "id": 152235,
      "category": "AGG. ASSAULT",
      "latitude": 39.33635,
      "longtitude": -76.61063399999999
    },
    {
      "id": 162902,
      "category": "LARCENY",
      "latitude": 39.3347214222,
      "longtitude": -76.61061881020001
    },
    {
      "id": 173995,
      "category": "AGG. ASSAULT",
      "latitude": 39.333647,
      "longtitude": -76.60994000000001
    },
    {
      "id": 175919,
      "category": "LARCENY",
      "latitude": 39.334653,
      "longtitude": -76.61061600000001
    },
    {
      "id": 177172,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.336929999999995,
      "longtitude": -76.61064300000001
    },
    {
      "id": 182193,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.333647,
      "longtitude": -76.60994000000001
    },
    {
      "id": 197990,
      "category": "BURGLARY",
      "latitude": 39.333702,
      "longtitude": -76.609951
    },
    {
      "id": 218487,
      "category": "BURGLARY",
      "latitude": 39.3373637707,
      "longtitude": -76.6106714693
    },
    {
      "id": 221412,
      "category": "LARCENY",
      "latitude": 39.334574009899995,
      "longtitude": -76.6106179512
    },
    {
      "id": 222365,
      "category": "BURGLARY",
      "latitude": 39.334653,
      "longtitude": -76.61061600000001
    },
    {
      "id": 231046,
      "category": "ROBBERY - STREET",
      "latitude": 39.3336545733,
      "longtitude": -76.6099738301
    },
    {
      "id": 240754,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.333751284499996,
      "longtitude": -76.6101788284
    },
    {
      "id": 253238,
      "category": "ROBBERY - STREET",
      "latitude": 39.33635,
      "longtitude": -76.61063399999999
    },
    {
      "id": 257423,
      "category": "LARCENY FROM AUTO",
      "latitude": 39.3336548153,
      "longtitude": -76.6099403947
    },
    {
      "id": 273852,
      "category": "BURGLARY",
      "latitude": 39.3348145348,
      "longtitude": -76.6106900606
    },
    {
      "id": 276535,
      "category": "ROBBERY - STREET",
      "latitude": 39.334780387100004,
      "longtitude": -76.6106191539
    },
    {
      "id": 278636,
      "category": "LARCENY",
      "latitude": 39.3357777946,
      "longtitude": -76.6106991499
    }
  ],
  "status_code": 200
}
```

