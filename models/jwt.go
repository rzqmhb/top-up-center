package models

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var JWTKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	jwt.StandardClaims
	Username string `json:"user_name"`
}