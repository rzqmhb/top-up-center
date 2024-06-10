package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rzqmhb/top-up-center/models"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		session_token, err := c.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				if c.Request.Header.Get("Content-Type") == "application/json" {
					c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
				return
				}
				c.Redirect(http.StatusSeeOther, "/login")
			}
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		claims := &models.Claims{}
		
		token, err := jwt.ParseWithClaims(session_token, claims, func(t *jwt.Token) (interface{}, error) {
			return models.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println(claims.Username, "\n", session_token, "\n", token)
		c.Set("username", claims.Username)
		c.Next()
	})
}