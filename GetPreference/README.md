# GetPreference:: Go 1.x REST API for LiveGood

**GetPreference** is a Go 1.x REST API for LiveGood.

## Features

* Give back user preference.


#### API Endpoint - POST
```URL
https://5cbj60afu4.execute-api.us-east-1.amazonaws.com/getPreference/
```

#### Request Body 

```JSON
{
  "username": "9088"
}
```

#### Result

```JSON
Response:
{
  "status": 200,
  "rent": 1000,
  "bedroom": 4,
  "bathroom": 4,
  "distance": 1,
  "crime": 0.6,
  "convenience": 0.4
}

```

