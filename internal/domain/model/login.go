package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID   int    `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

