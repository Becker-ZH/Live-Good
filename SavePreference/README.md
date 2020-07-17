# save_preference:: Go 1.x REST API for LiveGood

**save_preference** is a Go 1.x REST API for LiveGood.

## Features

* Update housing preference for a user.

### <update\>
The `update` method read `username` and all other related information, such as `rent`, `bedroom`, `bathroom`, and `distance`, and update in the row of that `username`.

#### API Endpoint - POST
```URL
https://h2p59lrq6k.execute-api.us-east-1.amazonaws.com/preference/
```

#### Request Body 

```JSON
{
  "username": "9088",
  "rent": 1000,
  "bedroom": 4,
  "bathroom": 4,
  "distance": 1,
  "crime": 0.6,
  "convenience": 0.4
}
```

#### Result

```JSON
Response:
{
  "status": 200,
  "feedback": "Update preference"
}
```

