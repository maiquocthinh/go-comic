package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

func (mm *middlewareManager) RequiredAuthJWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := utils.GetUserTokenFromHeader(&ctx.Request.Header)
		if err != nil {
			panic(common.NewUnauthorizedApiError(errors.New("Invalid token"), ""))
		}

		userClaims, err := utils.ParseJWTOfUser(tokenString, &mm.cfg.Server)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		inBlackList, err := mm.authRedisRepo.IsTokenInBlackList(ctx.Request.Context(), userClaims.ID)
		if err != nil {
			panic(common.NewInternalApiError(err, ""))
		}
		if inBlackList {
			panic(common.NewUnauthorizedApiError(err, ""))

		}

		ctx.Set(common.KeyUserClaims, userClaims)

		ctx.Next()
	}
}

func (mm *middlewareManager) OptionalAuthJWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := utils.GetUserTokenFromHeader(&ctx.Request.Header)
		if err != nil {
			ctx.Next()
			return
		}

		userClaims, err := utils.ParseJWTOfUser(tokenString, &mm.cfg.Server)
		if err != nil {
			ctx.Next()
			return
		}

		if inBlackList, err := mm.authRedisRepo.IsTokenInBlackList(ctx.Request.Context(), userClaims.ID); err != nil || inBlackList {
			ctx.Next()
			return
		}

		ctx.Set(common.KeyUserClaims, userClaims)

		ctx.Next()
	}
}
