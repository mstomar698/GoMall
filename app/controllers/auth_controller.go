package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mstomar698/GoMall/app/database"
	"github.com/mstomar698/GoMall/app/models"
	"github.com/mstomar698/GoMall/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "email or password is incorrect"
		check = false
	}
	return check, msg
}

var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")
var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for phone"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email or phone number already exist"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, err := utils.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating Token"})
		}
		user.Token = &token
		user.Refresh_token = &refreshToken
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, *&foundUser.User_id)
		utils.UpdateAllToken(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("token", token, 3600, "/", "", false, true)
		c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)
		c.Set("user", foundUser)
		c.Redirect(http.StatusFound, `/home/?={token}`)
		// c.JSON(http.StatusOK, foundUser)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
		c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
	}
}
