package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type jwtClaim struct {
	Id    int64
	Email string
	jwt.RegisteredClaims
}

func GenToken(email string, id int64, secret string) (string, error) {
	claim := &jwtClaim{
		id,
		email,
		jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			// 签发人
			Issuer: "yogen",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
}

func ParseToken(tokenString string, secret string) (email string, id int64, err error) {
	claim := new(jwtClaim)
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return
		// 校验token
	}
	if token.Valid {
		return claim.Email, claim.Id, nil
	}
	return "", 0, fmt.Errorf("token invalid!")
}
