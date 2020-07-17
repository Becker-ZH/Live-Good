# updateNearby:: Go 1.x REST API for LiveGood

**updateNearby** is a Go 1.x REST API for LiveGood.

## Features

* Using `jhuapartment`, `jhuplace`, and `aptplace` database tables to fetch nearby stores/restaurant for a particular apartment.

### <getReqID\>
The `getReqID` method read `latitude` and `longtitude` given from the request and find the corresponding apartment ID in `jhuapartment`.

### <getCombinedList\>
The `getCombinedList` method combines information of all nearby convenient use places for the apartment.

#### API Endpoint - POST
```URL
https://5p7kg3wvlj.execute-api.us-east-1.amazonaws.com/jhuNearby/
```

#### Request Body 

```JSON
{
  "type": "nearby",
  "latitude": 39.3357969,
  "longtitude": -76.6035516
}
```
- `type` is negligible.

#### Result

```JSON
{
  "type": "get_nearby",
  "all_nearby": [
    {
      "id": 1,
      "name": "Shun Lee Chinese Carry Out",
      "category": "restaurant",
      "address": "3405 Greenmount Ave, Baltimore",
      "latitude": 39.3299728,
      "longtitude": -76.60917959999999
    },
    {
      "id": 2,
      "name": "T & J’s Carryout",
      "category": "restaurant",
      "address": "3842 Crestlyn Rd, Baltimore",
      "latitude": 39.3356656,
      "longtitude": -76.5985864
    },
    {
      "id": 3,
      "name": "Mayflower Chinese Restaurant",
      "category": "restaurant",
      "address": "3407 Greenmount Ave, Baltimore",
      "latitude": 39.330051,
      "longtitude": -76.60903259999999
    },
    {
      "id": 4,
      "name": "York Food Market",
      "category": "market",
      "address": "532 Rose Hill Terrace, Baltimore",
      "latitude": 39.3388543,
      "longtitude": -76.6076975
    },
    {
      "id": 5,
      "name": "Crestlyn Market",
      "category": "market",
      "address": "3838 Crestlyn Rd, Baltimore",
      "latitude": 39.3355389,
      "longtitude": -76.5986542
    },
    {
      "id": 6,
      "name": "Belgian Food Market",
      "category": "market",
      "address": "841 Belgian Ave, Baltimore",
      "latitude": 39.340295,
      "longtitude": -76.603787
    },
    {
      "id": 7,
      "name": "Waverly Crabs",
      "category": "market",
      "address": "3400 Greenmount Ave, Baltimore",
      "latitude": 39.3300177,
      "longtitude": -76.6097407
    },
    {
      "id": 8,
      "name": "Waverly Tavern",
      "category": "bar",
      "address": "3801 Old York Rd, Baltimore",
      "latitude": 39.3356206,
      "longtitude": -76.6073189
    },
    {
      "id": 9,
      "name": "Mullan Park",
      "category": "park",
      "address": "4000 Old York Rd, Baltimore",
      "latitude": 39.3381252,
      "longtitude": -76.60850669999999
    },
    {
      "id": 10,
      "name": "Andover and North Hill Park",
      "category": "park",
      "address": "1116 Andover Rd, Baltimore",
      "latitude": 39.3369337,
      "longtitude": -76.5997192
    },
    {
      "id": 11,
      "name": "Chestnut Hill Park",
      "category": "park",
      "address": "601 Chestnut Hill Ave, Baltimore",
      "latitude": 39.3340291,
      "longtitude": -76.60615969999999
    },
    {
      "id": 12,
      "name": "JHU Guesthouse",
      "category": "hotel",
      "address": "509 E 38th St, Baltimore",
      "latitude": 39.3351568,
      "longtitude": -76.6091953
    }
  ],
  "status_code": 200
}
```

