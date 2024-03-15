package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"strings"
)

const (
	KeyUserClaims = "UserClaims"
)

func (mm *middlewareManager) AuthJWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerHeader := ctx.Request.Header.Get("Authorization")
		if bearerHeader == "" {
			panic(common.NewUnauthorizedApiError(errors.New("Invalid token"), ""))
		}
		headerParts := strings.Split(bearerHeader, " ")
		if len(headerParts) != 2 {
			panic(common.NewUnauthorizedApiError(errors.New("Invalid token"), ""))
		}

		tokenString := headerParts[1]

		userClaims, err := utils.ParseJWTOfUser(tokenString, &mm.cfg.Server)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		ctx.Set(KeyUserClaims, userClaims)

		ctx.Next()
	}
}
