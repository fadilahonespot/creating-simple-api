# Creating Simple Api
this application is a simple api application using golang and mysql
## go version
- 1.5
## library required
- gin (router) [link](https://github.com/gin-gonic/gin)
- gorm (orm)  [link](https://github.com/jinzhu/gorm)
- uuid [link](https://github.com/google/uuid)
- jwt (security) [link](https://github.com/dgrijalva/jwt-go)
- crypto (bycrip) [link](https://golang.org/x/crypto)

-- note: You don't need to install the libraries one by one, because in the project there is a go mod, when the project is run for the first time it will automatically download all the libraries needed 

## Installation
### Using Docker Makefile Command
- run terminal
```docker
 $~ make deploy
```
- view logs of running projects
```shell
$~ make logs
```
#### Local app access
- url
```local
localhost:7081
```
#### Access phpmyadmin
- url
```local
localhost:8085
```
- user login phpmyadmin
```access
user: root
password: pass123
```
### Without Docker
- change file .env.example name to .env
- Setting database mysql in .env file
- run in terminal
```terminal
  $~ go run main.go
```

## Testing
### Method POST : User Registered
```url
{{BASE_URL}}/user/register
```
- Request Body
```body
{
    "name": "andi",
    "email": "andi@example.com",
    "password": "123456"
}
```
### Methode POST : User login
```url
{{BASE_URL}}/user/login
```
- Request Body
```body
{
    "email": "andi@example.com",
    "password": "123456"
}
```
- Respon body
```respon
{
    "success": true,
    "message": "Success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiJDJhJDEwJEQuUS9oM3FTT2dVbXFPSGNvdjdCbS5VbEdUUkl2UC52MW13WjkuTnNGVjFqN3BRLm1VdnVHIiwidXNlcl9pZCI6MX0.pjVi1Gn1ddokTsok5EQPDO8d8n60JIiZq3rB2Pl_kVI"
    }
}
```
each account will get a different token code, you can copy the token code in the response body then enter it into the authorization section then select bearer token
![input bearer](https://github.com/fadilahonespot/creating-simple-api/raw/master/postman-bearer-token.PNG)

### Method POST : Add question
```url
{{BASE_URL}}/question
```
- Request Body
```json
{
    "question": "lorem ipsum dolor sit amet consectetur adipiscing elit"
}
```
### Method GET : Get all question
```url
{{BASE_URL}}/question?page=1&limit=20
```
### Method GET : Get detail question
```url
{{BASE_URL}}/question/{{UUID}}
```
### Method PUT : Update question
```url
{{BASE_URL}}/question/{{UUID}}
```
- request body
```json
{
    "question": "INI UNTUK UPDATE 3",
    "is_active": true
}
```
### Methode DELETE : Delete question
```url
{{BASE_URL}}/question/{{UUID}}
```


