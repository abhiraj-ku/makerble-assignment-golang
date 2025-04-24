package auth

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/abhiraj-ku/health_app/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// jwtSecret := config.AppConfig.JWTSecret

func GenerateToken(UserID int64, Role string) (string, error) {
	expirationTime, err := time.ParseDuration(config.AppConfig.ExpirationTime)
	log.Println(expirationTime)
	if err != nil {
		slog.Error("failed to parse the expiration time from config", fmt.Sprint("error"), err)
		return "", fmt.Errorf("error parsing expiration time from env: %v", err)
	}
	claims := &Claims{
		UserID: UserID,
		Role:   Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println(token)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
