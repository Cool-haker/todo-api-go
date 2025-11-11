package middleware

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	jwtSecret []byte
	once      sync.Once
)

func LoadJWTSecret() []byte {
	once.Do(func() {
		if err := godotenv.Load("internal/db/.env"); err != nil {
			log.Fatalf("error loading .env: %v", err)
		}

		key := os.Getenv("JWT_SECRET")
		if key == "" {
			log.Fatal("SECRET_KEY not set in .env")
		}

		jwtSecret = []byte(key)
	})

	return jwtSecret
}

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, fmt.Errorf("token expired")
		}
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, jwt.ErrInvalidKey
		}
		return uint(userIDFloat), nil
	}

	return 0, jwt.ErrInvalidKey
}
