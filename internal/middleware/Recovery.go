package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				var apiErr *common.ApiError

				apiErr, ok := err.(*common.ApiError) // parse to ApiError
				if !ok {
					apiErr = common.NewInternalApiError(err.(error), "Internal Server Error")
				}

				ctx.AbortWithStatusJSON(apiErr.StatusCode, apiErr)
				panic(err)
			}

		}()

		ctx.Next()
	}
}
