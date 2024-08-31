package customauth

import (
	"HangAroundBackend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTsecret = []byte("secret_token_from_env")

// signing jwt token with username and email
// Could be modified
type AccessTokenClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(id uint, email, username, role string) (string, error) {
	expirationTime := time.Now().Add(4 * time.Hour)

	claims := &AccessTokenClaims{
		ID:       id,
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTsecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type RefreshTokenClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(id uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := RefreshTokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTsecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func init() {
	if config.GetEnv("AUTH_JWT_SECRET") != "" {
		JWTsecret = []byte(config.GetEnv("AUTH_JWT_SECRET"))
	} else {
		panic("AUTH_JWT_SECRET is not set")
	}
}
