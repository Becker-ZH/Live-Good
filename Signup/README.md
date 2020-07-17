# Signup:: Go 1.x REST API for LiveGood

**Signup** is a Go 1.x REST API for LiveGood.

## Features

* Sign up user information into the `user_info` table with `bcrypt`.

### <save\>
The `save` method receives a request containing user information, such as `first_name`, `last_name`, `user_name`, `password` and stores the corresponding user information into the `user_info` db table with `password` hashed with `bcrypt`. If the user already exists, then it will return a notified message with status code equals `404`; if success, then return status code `200`.

#### API Endpoint - POST
```URL
https://8po7546c75.execute-api.us-east-1.amazonaws.com/signup/
```

#### Request Body 

```JSON
{
  "first_name": "jiya",
  "last_name": "zhang",
  "user_name": "9088",
  "password": "9088"
}
```

#### Result

```JSON
Response:
{
  "status": 200,
  "feedback": "Sign up success"
}
{
  "status": 404,
  "feedback": "User name already exists"
}
```

