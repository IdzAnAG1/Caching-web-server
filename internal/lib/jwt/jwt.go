package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserLogin    string
	UserPassword string
	jwt.RegisteredClaims
}

func SignToken(login, serverToken, password string, TTL time.Duration) (string, error) {
	// Todo This shit returned the error, solved it | hmmm solved ?
	exptime := time.Now().Add(TTL)
	claims := &Claims{
		UserLogin:    login,
		UserPassword: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exptime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(serverToken))
}
