package middleware

import (
	"net/http"
	"os"
	"strings"
	"worker-management/internal/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(*model.Claims)
		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.UserType)
		c.Next()
	}
}

func GetUserID(c *gin.Context) int {
	v, _ := c.Get("user_id")
	return v.(int)
}

func GetUserType(c *gin.Context) string {
	v, _ := c.Get("user_type")
	return v.(string)
}

func RequireRole(expected string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := GetUserType(c)
		if role != expected {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized: only " + expected})
			return
		}
		c.Next()
	}
}
