package request

import (
	"github.com/dgrijalva/jwt-go"
)
type Claims struct {
	Email 			string 		`json:"email"`
	Id 				int64 		`json:"id"`
	jwt.StandardClaims
}
