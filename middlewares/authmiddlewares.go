package middlewares

import (
	"example/Go/initializers"
	"example/Go/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectRoute(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		
		return []byte(os.Getenv("JWT_SECRETE")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()
		
	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
}

func RestrictedRoute(role string) gin.HandlerFunc {
	return func(c *gin.Context) {		
		user, _ := c.Get("user")		

		
		if user.(models.User).Role != role {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to perform this action"})
			return
		}		
		c.Next()
	}
}