package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	InvalidTokenStringError = "token is invalid"
)

var (
	UndefinedTokenError = fmt.Errorf("undefined token")
	InvalidTokenError   = fmt.Errorf(InvalidTokenStringError)
)

type ServiceJWT struct {
	privetKey      *rsa.PrivateKey
	publicKey      *rsa.PublicKey
	RefreshTimeExp time.Duration
	AccessTimeExp  time.Duration
}

func NewServiceJWT(privetKey *rsa.PrivateKey, publicKey *rsa.PublicKey,
	refreshTimeExp time.Duration, accessTimeExp time.Duration) *ServiceJWT {
	return &ServiceJWT{
		privetKey:      privetKey,
		publicKey:      publicKey,
		RefreshTimeExp: refreshTimeExp,
		AccessTimeExp:  accessTimeExp,
	}
}

func (j *ServiceJWT) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(j.privetKey)
	if err != nil {
		return "", fmt.Errorf("failed create token: %v", err)
	}

	return tokenString, nil
}

func (j *ServiceJWT) DecodeKey(tokenString string) (*jwt.RegisteredClaims, error) {
	if tokenString == "" {
		return nil, UndefinedTokenError
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.publicKey, nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, InvalidTokenError
}
