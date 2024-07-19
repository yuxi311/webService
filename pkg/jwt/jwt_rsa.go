package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const ExpireTimeHours = 24 * time.Hour

type Token struct {
	jwt.StandardClaims
	Role     int32  `json:"role,omitempty"`
	Username string `json:"username,omitempty"`
}

func CreateToken(username string, role int32) (string, error) {
	tk := &Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireTimeHours).Unix(),
		},
		Role:     role,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tk)

	signingKey, err := GetPrivateKey()
	if err != nil {
		return "", err
	}
	return token.SignedString(signingKey)
}

func ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		pubKey, err := GetPublicKey()
		return pubKey, err
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
