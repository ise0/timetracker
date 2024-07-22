package api

import (
	"time"
	"timetracker/src/lib"

	"github.com/gin-gonic/gin"
)

func loggerMiddleware(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	lib.Logger.Debugw("http request",
		"method", ctx.Request.Method,
		"date", start,
		"statusCode", ctx.Writer.Status(),
		"latency", time.Since(start),
		"errmsg", ctx.Errors.ByType(1<<0).String(),
		"clientIP", ctx.ClientIP(),
		"path", path,
	)
}
