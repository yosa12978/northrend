package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/northrend/services"
)

func Logger() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		startTime := time.Now().UnixNano()
		ctx.Next()
		latency := time.Now().UnixNano() - startTime
		logger := services.NewConsoleLogger("request")
		logger.Fields(map[string]interface{}{
			"method":   ctx.Request.Method,
			"endpoint": ctx.Request.RequestURI,
			"code":     ctx.Writer.Status(),
			"latency":  latency,
		})
	}
}
