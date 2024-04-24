package middleware

import (
	"net/http"
	"os"
	"time"
	"fmt"

	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//Get cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode validae it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check the expire
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach
		c.Set("user", user)

		//continue

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// RequireRole checks if the authenticated user has the required role.
func RequireRole(requiredRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the context
		user, exists := c.Get("user")

		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Type assert to the User model
		u, ok := user.(models.User)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the user has the required role
		if u.Role != requiredRole {
			c.AbortWithStatus(http.StatusForbidden) // 403 Forbidden
			return
		}

		// If the user has the correct role, continue to the next handler
		c.Next()
	}
}
