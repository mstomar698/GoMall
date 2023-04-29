package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "home", gin.H{
			"title": "",
		})
	}
}

func HomeSecured() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userCollection.Find(context.Background(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var userList []models.User
		for users.Next(context.Background()) {
			var user models.User
			if err := users.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			userList = append(userList, user)
		}
		userName := c.GetString("first_name")
		userType := c.GetString("user_type")
		userEmail := c.GetString("email")
		userLocation := c.GetString("first_name")
		userPhone := c.GetString("first_name")
		userTitle := c.GetString("last_name")
		userId := c.GetString("uid")
		admin := c.GetBool("admin")
		c.HTML(http.StatusOK, "homesecured", gin.H{
			"username":  userName,
			"user_type": userType,
			"admin":     admin,
			"email":     userEmail,
			"phone":     userPhone,
			"location":  userLocation,
			"title":     userTitle,
			"user_id":   userId,
			"users":     userList,
		})
	}
}

func LoginScreen() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "login", gin.H{
			"title": "",
		})
	}
}

func SignupScreen() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "signup", gin.H{
			"title": "",
		})
	}
}

func EditUserScreen() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		c.HTML(200, "edituser", gin.H{
			"user_id": userId,
		})
	}
}
func DeleteUserScreen() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		c.HTML(200, "deleteuser", gin.H{
			"user_id": userId,
		})
	}
}
func CreateUserScreen() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "createuser", gin.H{
			"title": "",
		})
	}
}
