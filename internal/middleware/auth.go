package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"strings"
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

		ctx.Set(common.KeyUserClaims, userClaims)

		ctx.Next()
	}
}

func (mm *middlewareManager) VerifyJWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, err := utils.GetUserTokenClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		if inBlackList, err := mm.authRedisRepo.IsTokenInBlackList(ctx.Request.Context(), userClaims.ID); err != nil || inBlackList {
			panic(common.NewUnauthorizedApiError(errors.New("Invalid token"), ""))
		}

		ctx.Next()
	}
}
