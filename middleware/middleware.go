package middleware

import (
	"fmt"
	"net/http"
	"project-inventory-book/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthValid(c *gin.Context) {
	var tokenString string
	tokenString = c.Query("auth")
	if tokenString == "" {
		tokenString = c.PostForm("auth")
		if tokenString == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"content": "token nil",
			})
			c.Abort()
		}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
		}
		return []byte(models.SECRET), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"content": "token is expiry",
		})
		c.Abort()
	}
}
