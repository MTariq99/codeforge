package models

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserId    int64  `json:"user_id" db:"user_id"`
	UserName  string `json:"user_name" db:"user_name"`
	UserEmail string `json:"user_email" db:"user_email"`
	jwt.StandardClaims
}
