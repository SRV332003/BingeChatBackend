package customauth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyAccessToken(accessToken string) (*AccessTokenClaims, error) {
	claims :=
		&AccessTokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid Token")
	}
	return claims, nil
}

func VerifyRefreshToken(refreshtoken string) (*RefreshTokenClaims, error) {
	claims :=
		&RefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(refreshtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid Token")
	}
	return claims, nil
}
