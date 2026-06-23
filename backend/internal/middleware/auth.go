package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Tajny klucz dla JWT
var JwtKey = []byte("super-secret-uni-key")

// RequireAuth sprawdza token
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wymagana autoryzacja"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowy format tokenu"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["user_id"])
			c.Set("userRole", claims["role"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieważny token", "details": err.Error()})
			c.Abort()
		}
	}
}

// RequireAdmin pozwala na dostęp adminowi oraz superadminowi
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || (role != "admin" && role != "superadmin") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Dostęp dozwolony tylko dla administratorów"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireSuperAdmin pozwala na dostęp wyłącznie superadminowi
func RequireSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "superadmin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Dostęp dozwolony tylko dla głównego administratora (SuperAdmin)"})
			c.Abort()
			return
		}
		c.Next()
	}
}
