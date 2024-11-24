package middleware

import (
	"i-shop/controllers"
	"i-shop/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AutoMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
	
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Println("[ERROR] Authorization header missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			if err.Error() == "token is expired" {
				log.Printf("[ERROR] Token expired: %v\n", err)
				controllers.HandleResponse(c, http.StatusUnauthorized, "Token expired")
			} else {
				log.Printf("[ERROR] Invalid token: %v\n", err)
				controllers.HandleResponse(c, http.StatusUnauthorized, "Invalid token")
			}
			c.Abort()
			return
		}

		if claims.Role != role {
			log.Printf("[ERROR] Access denied for role: %s, required: %s\n", claims.Role, role)
			controllers.HandleResponse(c, http.StatusForbidden, "Access denied")
			c.Abort()
			return
		}

		log.Printf("[INFO] Authorized request: Email: %s, Role: %s\n", claims.Email, claims.Role)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}
