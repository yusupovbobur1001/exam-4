package middleware

import (
	"api_service/api/token"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("authorization header is required"))
			return
		}
		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("invalid token: %s", err))
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("invalid token claims: %s", err))
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		jwtClaims := claims.(jwt.MapClaims)
		sub := jwtClaims["role"].(string)

		obj := c.FullPath()
		act := c.Request.Method
		fmt.Println(sub, obj, act)
		allowed, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during authorization"})
			return
		}

		if !allowed {
			fmt.Println("-0-0-0-0-00-0-0-0-0-0-0-0-0-")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
