package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtAccessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var jwtRefreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GenerateTokens(username, role string) (Tokens, error) {
	atClaims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}

	rtClaims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	accessToken, err := at.SignedString(jwtAccessSecret)
	if err != nil {
		return Tokens{}, err
	}

	refreshToken, err := rt.SignedString(jwtRefreshSecret)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func ValidateToken(tokenStr string, isRefresh bool) (*Claims, error) {
	var secret = jwtAccessSecret
	if isRefresh {
		secret = jwtRefreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
