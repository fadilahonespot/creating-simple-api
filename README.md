# Creating Simple Api
## go version
- 1.5
## library required
- gin (router)[a link](https://github.com/gin-gonic/gin)
- gorm (orm)  [a link](https://github.com/jinzhu/gorm)
- uuid [a link](github.com/google/uuid)
- jwt (security) [a link](https://github.com/google/uuid)
- crypto (bycrip) [a link](https://golang.org/x/crypto)
## Environment Variabel Declaration
- 1. Define .env file
```env
mv .env
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
SECRET = "yourKey"
```
-- note: this is a tool for generating passwords, you can make the password as you like
