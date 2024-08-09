# User API Spec

## Login
Endpoint : POST /api/users/login
Requset body :
````json
{
    "username" : "john doe",
    "password" : "secret"
}
````

Response body (success) :
````json
{
    "statusCode" : 200,
    "data" : {
        "username" : "John Doe",
        "password"  : "secret",
        "token" : "jwt-token"
    }
}
````

Response body (error) :
````json
{
    "statusCode" : 4xx,
    "errors"    : "USERNAME_PASSWORD_UNMATCH"
}
````



## Register
Endpoint : POST /api/users/
Requset body :
````json
{
    "username" : "john doe",
    "password" : "secret"
}
````

Response body (success) :
````json
{
    "statusCode" : 200,
    "data" : {
        "username" : "John Doe",
        "password"  : "secret"
    }
}
````

Response body (error) :
````json
{
    "statusCode" : 4xx,
    "errors"    : "USERNAME_ALREADY_REGISTERED"
}
````




## GetUser

Endpoint : GET /api/users/current
Header :
- authorization : token

Response body (success) :
````json
{
    "statusCode" : 200,
    "data" : {
        "username" : "John Doe",
        "password"  : "secret"
    }
}
````

Response body (error) :
````json
{
    "statusCode" : 4xx,
    "errors"    : "UNAUTHORIZED"
}
````

## Update User
Endpoint : PATCH /api/users/
- PUT : update all columns
- PATCH : only update some column(s)

Header :
- authorization : token

Requset body :
````json
{
    "username" : "john doe", // optional
    "password" : "secret" // optional
}
````

Response body (success) :
````json
{
    "statusCode" : 200,
    "data" : {
        "username" : "John Doe",
        "password"  : "secret"
    }
}
````

Response body (error) :
````json
{
    "statusCode" : 4xx,
    "errors"    : "UNAUTHORIZED"
}
````



## GetUser

Endpoint : DELETE /api/users/current
Header :
- authorization : token

Response body (success) :
````json
{
    "statusCode" : 200,
    "data" : true
}
````

Response body (error) :
````json
{
    "statusCode" : 4xx,
    "errors"    : "UNAUTHORIZED"
}
````






## Forgot Password
## Logout