package userRouter

import (
	"github.com/gin-gonic/gin"
)

func Apply(g *gin.RouterGroup) {
	g.GET("/info", getUserByPassport)
	g.POST("/user", addUser)
	g.PATCH("/user", updateUser)
	g.DELETE("/user", deleteUser)
	g.GET("/user-stat", getUserStat)
	g.GET("/users", getUsers)
}
