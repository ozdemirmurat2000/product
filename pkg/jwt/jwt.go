package jwt

import (
	"errors"
	"productApp/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int, username string, isAdmin bool, expiresAt time.Time) (string, error) {
	claims := &Claims{
		UserID:   userId,
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 33 * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JWTSecretKey))
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecretKey), nil
	})

	if err != nil {
		return nil, errors.New("geçersiz token")
	}
	if !token.Valid {
		return nil, errors.New("token geçersiz")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("token claim hatası")
	}

	return claims, nil
}

func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return GenerateJWT(claims.UserID, claims.Username, claims.IsAdmin, time.Now().Add(time.Hour*7*24))
}
