# INFO

### Project Structure

```scss
backend
    |__ cmd
    |     |__ main.go
    |     |__ .env
    |
    |__ app
    |      |__ app.go
    |      |
    |      |__controller
    |      |           |__ auth_controller.go
    |      |           |__ temp_controller.go
    |      |           |__ user_controller.go
    |      |__database
    |      |         |__ db.go
    |      |         |__ env.go
    |      |__models
    |      |       |__ user_model.go
    |      |__routes
    |      |       |__ auth.go
    |      |       |__ home.go
    |      |       |__ user.go
    |      |       |__ screen.go
    |      |__utils
    |      |      |__ authenticate.go
    |      |      |__ token.go
    |______|______|__ user_type.go
    |
    |__ web
    |     |__ template
    |     |      |__ home.html
    |     |      |__ login.html
    |     |      |__ signup.html
    |     |      |__ homesecured.html
    |     |      |__ allusers.html
    |     |      |__ createuser.html
    |     |      |__ deleteuser.html
    |     |      |__ ediuser.html
    |     |      |__ edit.html
    |     |__ static
    |     |      |__ css
    |     |      |      |__ home.css
    |     |      |      |__ login.css
    |     |      |      |__ signup.css
    |     |      |__ js
    |     |      |      |__ index.js
    |     |      |__ img
    |     |      |      |__ B-rem.png
    |     |______|______|__ favicon.ico
    |
    |__ vendor # this folder is created by go mod vendor
    |__ test # this folder is created for test cases of the project
    |__ go.mod
    |__ go.sum
    |__ README.md
    |__ INFO.md
```

> Work by [Mayank](http.//github.com/mstomar698)


> go get github.com/dgrijalva/jwt-go github.com/gin-gonic/gin github.com/go-playground/validator/v10 github.com/joho/godotenv go.mongodb.org/mongo-driver golang.org/x/crypto