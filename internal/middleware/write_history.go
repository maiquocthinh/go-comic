package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"strconv"
	"time"
)

func (mm *middlewareManager) WriteHistory() gin.HandlerFunc {
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

		comicID, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			ctx.Next()
			return
		}
		chapterID, err := strconv.Atoi(ctx.Param("chapterID"))
		if err != nil {
			ctx.Next()
			return
		}

		now := time.Now()

		jsonData, err := json.Marshal(&models.HistoryView{
			UserID:    userClaims.UserID,
			ComicID:   comicID,
			ChapterID: chapterID,
			ViewedAt:  &now,
		})

		err = mm.pubsub.Publish(ctx.Request.Context(), common.TopicWriteHistoryView, string(jsonData))
		if err != nil {
			fmt.Println(err)
			ctx.Next()
			return
		}

		ctx.Next()
	}
}
