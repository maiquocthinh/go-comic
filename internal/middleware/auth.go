package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

func (mm *middlewareManager) AuthJWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := utils.GetUserTokenFromHeader(&ctx.Request.Header)
		if err != nil {
			panic(common.NewUnauthorizedApiError(errors.New("Invalid token"), ""))
		}

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
		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		if inBlackList, err := mm.authRedisRepo.IsTokenInBlackList(ctx.Request.Context(), userClaims.ID); err != nil || inBlackList {
			panic(common.NewUnauthorizedApiError(errors.New("Invalid token"), ""))
		}

		ctx.Next()
	}
}
