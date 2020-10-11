package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
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
	Secret []byte // string
	Issuer string
}

func MakeJWTHS265(secret string, issuer string) *JWTHS265 {
	return &JWTHS265{
		Secret: []byte(secret),
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
	tokenStringified, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.Secret)
	if err != nil {
		return "", 0, err
	}
	return tokenStringified, expirationTime.Unix(), nil

}

// Validates the Token
func (j *JWTHS265) ValidateWebToken(token string) (isValid bool, jwtStdClaims *Claims) {
	var claim Claims
	tkn, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
	}
	if !tkn.Valid {
		return false, nil
	}
	return true, &claim
}
