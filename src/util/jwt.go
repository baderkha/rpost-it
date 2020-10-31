package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IJwtHelper interface {
	GenerateWebToken(subject string, ttl int64) (token string, validity int64, err error)
	ValidateWebToken(token string) (isValid bool, jwtStdClaims *Claims)
}

// JWT CLAIMS
type Claims struct {
	jwt.StandardClaims
}

// HMAC 265 TOKEN
type JWTHS265 struct {
	Secret string // string
	Issuer string
}

func MakeJWTHS265(secret string, issuer string) *JWTHS265 {
	return &JWTHS265{
		Secret: secret,
		Issuer: issuer,
	}
}

// Generates a JWT token
func (j *JWTHS265) GenerateWebToken(subject string, ttlMinutes int64) (token string, validity int64, err error) {
	expirationTime := time.Now().Add(time.Duration(ttlMinutes) * time.Minute)
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    j.Issuer,
			Subject:   subject,
		},
	}

	tokenStringified, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(j.Secret))

	if err != nil {
		fmt.Println(err.Error())
		return "", 0, err
	}
	return tokenStringified, expirationTime.Unix(), nil

}

// Validates the Token
func (j *JWTHS265) ValidateWebToken(token string) (isValid bool, jwtStdClaims *Claims) {
	fmt.Println(token)
	tkn, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(j)
		return []byte(j.Secret), nil
	})

	if err != nil {
		return false, nil
	}

	if claims, ok := tkn.Claims.(*Claims); ok && tkn.Valid {
		return true, claims
	}

	return false, nil

}
