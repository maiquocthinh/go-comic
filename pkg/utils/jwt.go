package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"time"
)

const (
	jtiPrefix = "jti:"
)

type UserTokenClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	jwt.RegisteredClaims
}

func GenerateJWTForUser(user *entities.User, cfg *config.ServerConfig) (string, error) {
	now := time.Now()
	expirationTime := time.Duration(cfg.JwtUserExpiration) * time.Second

	claims := UserTokenClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jtiPrefix + uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(expirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString([]byte(cfg.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}
