package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	sign []byte
}

func NewJWT(sign string) *JWT {
	return &JWT{
		sign: []byte(sign),
	}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}

func (j *JWT) GenerateJwt(UserID int, days int, TokenType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * time.Duration(days)).Unix(),
		},
		UserID,
		TokenType,
	})

	tokenString, err := token.SignedString(j.sign)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}
		return j.sign, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, ErrExpired
	}
	return claims, nil
}
