package userRouter

import (
	"strconv"
	"timetracker/src/lib"
	userModel "timetracker/src/model/user"

	"github.com/gin-gonic/gin"
)

type getUserStatRes struct {
	TaskId    int    `json:"taskId"`
	TaskName  string `json:"name"`
	SpentTime string `json:"spentTime"`
	Completed bool   `json:"completed"`
}

func getUserStat(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.JSON(400, "provide valid userId")
		return
	}

	r, err := userModel.GetUserStat(userId)
	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	var res = make([]getUserStatRes, len(r))
	for i, v := range r {
		res[i] = getUserStatRes{v.TaskId, v.TaskName, v.SpentTime.String(), v.Completed}
	}

	c.JSON(200, res)
}
