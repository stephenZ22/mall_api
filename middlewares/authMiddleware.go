package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization token not provided",
			})

			c.Abort()
			return
		}

		secretKey := os.Getenv("SecretKey")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization token",
			})

			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization token",
			})

			c.Abort()
			return
		}

		c.Set("login_name", claims["login_name"])
		c.Next()
	}

}
