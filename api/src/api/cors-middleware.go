package api

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func corsMiddleware(ctx *gin.Context) {
	cors.New(cors.Options{AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), " "),
		AllowCredentials: true, AllowedMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"}}).HandlerFunc(ctx.Writer, ctx.Request)
}
