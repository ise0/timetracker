package userRouter

import (
	"strconv"
	"timetracker/src/lib"
	userModel "timetracker/src/model/user"

	"github.com/gin-gonic/gin"
)

func deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.JSON(400, "provide valid userId")
		return
	}

	err = userModel.DeleteUser(userId)
	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	c.JSON(200, "success")
}
