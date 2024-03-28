package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"strconv"
)

func (mm middlewareManager) IncreaseView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		jsonData, err := json.Marshal(&models.IncreaseView{
			ComicID:   comicID,
			ChapterID: chapterID,
		})

		err = mm.pubsub.Publish(ctx.Request.Context(), common.TopicIncreaseView, string(jsonData))
		if err != nil {
			fmt.Println(err)
			ctx.Next()
			return
		}

		ctx.Next()
	}
}
