package utils

import (
	"banners_service/internal/models"
	"banners_service/pkg/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func JwtCreate(user models.User) (string, error) {

    config := config.GetConfig()
    now := time.Now().UTC()
    tokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub":  user.ID,
        "role": user.Role,
        "exp":  now.Add(config.JwtExpiresIn).Unix(),
        "iat":  now.Unix(),
        "nbf":  now.Unix(),
    })

    tokenString, err := tokenByte.SignedString([]byte(config.JwtSecret))
    if err != nil {
		return "", fmt.Errorf("generating JWT token failed: %v", err)
    }
	return tokenString, nil
}