package api

import (
	taskRouter "timetracker/src/api/routes/task"
	userRouter "timetracker/src/api/routes/user"

	"github.com/gin-gonic/gin"
)

var Engine = gin.New()

func init() {
	Engine.Use(gin.Recovery(), loggerMiddleware, corsMiddleware)

	api := Engine.Group("api")

	userRouter.Apply(api)
	taskRouter.Apply(api)
}
