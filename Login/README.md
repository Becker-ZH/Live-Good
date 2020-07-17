# Login:: Go 1.x REST API for LiveGood

**Login** is a Go 1.x REST API for LiveGood.

## Features

* Log in API.

### <save\>
The `read` method read `user_name` and unhashed `password` to compare with hashed password stored into `user_info`.

#### API Endpoint - POST
```URL
https://dr42jeysxa.execute-api.us-east-1.amazonaws.com/login/
```

#### Request Body 

```JSON
{
  "user_name": "9088",
  "password": "9088"
}
```

#### Result

```JSON
Response:
{
  "status": 200,
  "feedback": "Login success"
}
```

