package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const ExpireTimeHours = 24 * time.Hour

type Token struct {
	jwt.StandardClaims
	TokenHolderRole int32
	TokenHolderName string `json:"username,omitempty"`
}

func CreateToken(username string, role int32) (string, error) {
	tk := &Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireTimeHours).Unix(),
		},
		TokenHolderRole: role, 
		TokenHolderName: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tk)

	signingKey, err := GetPrivateKey()
	if err != nil {
		return "", err
	}
	return token.SignedString(signingKey)
}
