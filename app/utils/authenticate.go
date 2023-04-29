package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cr := ""
		HeaderToken := c.Request.Header.Get("token")
		if HeaderToken == "" {
			cr = c.Request.Header.Get("cookie")
			cookies := strings.Split(cr, "; ")
			CookieToken := ""
			for _, cookie := range cookies {
				if strings.HasPrefix(cookie, "token=") {
					CookieToken = strings.TrimPrefix(cookie, "token=")
					break
				}
			}
			if CookieToken == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
				c.Abort()
				return
			}
			cr = CookieToken
		} else {
			cr = HeaderToken
		}
		claims, err := ValidateToken(cr)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		Admin := c.GetString("user_type")
		if Admin == "ADMIN" {
			c.Set("admin", true)
		} else {
			c.Set("admin", false)
		}
		c.Next()
	}
}
