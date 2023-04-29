package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUserType() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType := c.GetString("user_type")
		fmt.Println(userType)
		if userType == "ADMIN" {
			return
		} else {
			c.Abort()
			c.JSON(401, gin.H{"error": "Unauthorized to access this resource"})
		}
		c.Next()
	}
}

// func MatchUserTypeToUserid(c *gin.Context, userId string) (err error) {
// 	userType := c.GetString("user_type")
// 	uid := c.GetString("uid")
// 	err = nil

// 	if userType == "USER" && uid != userId { //to check if he is a user and he is using his id
// 		err = errors.New("Unauthorized to access this resource")
// 	}
// 	err = CheckUserType(c, userType)
// 	return err
// }
