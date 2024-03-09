package middleware

import (
	"bytes"
	"fmt"
	"log"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

const (
	errorLogFileBase = "./logs/error/"
)

var logger = logrus.New()
var today time.Time = time.Now()

func init() {
	logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] \n[MESSAGE]:\n\t%msg% \n[ERROR]:\n\t%error% \n[HEADERS]:\n\t%headers% \n[WHERE]:\n\t%where% \n\n\n",
	})

	if _, err := os.Stat(errorLogFileBase); os.IsNotExist(err) {
		os.MkdirAll(errorLogFileBase, os.ModePerm)
	}
	loadFileToLogger()
}

func ErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !dateEqual(today, time.Now()) {
			today = time.Now()
			loadFileToLogger()
		}

		defer func() {
			if err := recover(); err != nil {
				// get headers
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					// hidden token in header
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}

					headers[idx] = "\t" + headers[idx]
				}
				headersToStr := strings.Join(headers, "\r\n")

				// get where error panic
				where := new(bytes.Buffer)
				pc, _, _, ok := runtime.Caller(4)
				if ok {
					fn := runtime.FuncForPC(pc)
					file, line := fn.FileLine(pc)
					fmt.Fprintf(where, "%s:%d (0x%x)", file, line, fn.Entry())
				}

				// parse error
				var errCause, errMsg string
				if apiErr, ok := err.(*common.ApiError); ok {
					errCause = apiErr.Error()
					errMsg = apiErr.Message
				} else {
					_err := err.(error)
					errCause = _err.Error()
					errMsg = "Internal error"
				}

				logger.WithFields(logrus.Fields{
					"error":   errCause,
					"where":   string(where.Bytes()),
					"headers": strings.TrimSpace(headersToStr),
				}).Error(errMsg)

				panic(err)
			}
		}()

		ctx.Next()
	}
}

func loadFileToLogger() {
	dailyLogFile := errorLogFileBase + formatTime(today) + ".log"
	f, err := os.OpenFile(dailyLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}

	logger.SetOutput(f)
}

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func formatTime(date time.Time) string {
	return date.Format("2006-01-02")
}
