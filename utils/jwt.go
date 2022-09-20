package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/wonderivan/logger"
)

var JWTToken jwtToken

type jwtToken struct{}

type CostomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

const (
	SECRET = "admin"
)

func (*jwtToken) ParseToken(tokenString string) (claims *CostomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CostomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		logger.Error("ParseToken", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("TokenMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("TokenExpired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("TokenNotValidYet")
			} else {
				return nil, errors.New("TokenMalformed")
			}
		}
	}
	// 转换*CostomClaims类型并返回
	if claims, ok := token.Claims.(*CostomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("TokenMalformed")
}
