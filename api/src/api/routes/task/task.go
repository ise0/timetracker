package taskRouter

import "github.com/gin-gonic/gin"

func Apply(g *gin.RouterGroup) {
	g.POST("/task", addTask)
	g.PATCH("/task", completeTask)
}
