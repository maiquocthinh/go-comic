package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"net/http"
	"strings"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	tokenSigned, err := token.SignedString([]byte(cfg.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}

func ParseJWTOfUser(tokenString string, cfg *config.ServerConfig) (*UserTokenClaims, error) {
	var claims UserTokenClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return &claims, nil
}

func GetUserClaimsFromContext(ctx *gin.Context) (*UserTokenClaims, error) {
	claims, exists := ctx.Get(common.KeyUserClaims)
	if !exists {
		return nil, errors.New("Invalid token")
	}

	userClaims, ok := claims.(*UserTokenClaims)
	if !ok {
		return nil, errors.New("Invalid token")
	}

	return userClaims, nil
}

func GetUserTokenFromHeader(header *http.Header) (string, error) {
	bearerHeader := header.Get("Authorization")
	if bearerHeader == "" {
		return "", errors.New("Cannot extrac token.")
	}
	headerParts := strings.Split(bearerHeader, " ")
	if len(headerParts) != 2 {
		return "", errors.New("Cannot extrac token.")
	}

	tokenString := headerParts[1]

	return tokenString, nil
}
