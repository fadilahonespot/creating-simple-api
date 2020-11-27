# Creating Simple Api
this application is a simple api application using golang and mysql
## go version
- 1.5
## library required
- gin (router) [link](https://github.com/gin-gonic/gin)
- gorm (orm)  [link](https://github.com/jinzhu/gorm)
- uuid [link](github.com/google/uuid)
- jwt (security) [link](https://github.com/google/uuid)
- crypto (bycrip) [link](https://golang.org/x/crypto)
## Environment Variabel Declaration
- 1. Define .env file
```env
 $~ touch .env
 $~ nano .env
```
- 2. Define PORT in .env file, for example:
```port
PORT = 7081
```
- 3. Define MYSQL in .env file, for example:
```mysql
MYSQL = root:@tcp(127.0.0.1:3306)/creating_simple_api?parseTime=true
```
-- note: creating_simple_api is the name of the database, match the database that was created previously
- 4. Define SECRET in .env file, for example:
```secret
SECRET = yourKey
```
-- note: this is a tool for generating passwords, you can make the password as you like
## Installation
### Using Docker
```docker
 $~ docker-compose up --build -d
```
### Without Docker
```terminal
  $~ go run ./main.go
```

## Testing
- Method POST : User Registered
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
- Methode POST : User login
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

- Method POST : add question
```url
{{BASE_URL}}/question
```
- Request body add question
```json
{
    "question": "lorem ipsum dolor sit amet consectetur adipiscing elit"
}
```
- Method GET : Get all question
```url
{{BASE_URL}}/question?page=1&limit=20
```
- Method GET : Get detail question
```url
{{BASE_URL}}/question/{{UUID}}
```
- Method PUT : Update question
```url
{{BASE_URL}}/question/{{UUID}}
```
- request body update question
```json
{
    "question": "INI UNTUK UPDATE 3",
    "is_active": true
}
```
- Methode DELETE : delete question
```url
{{BASE_URL}}/question/{{UUID}}
```


