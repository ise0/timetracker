package taskRouter

import (
	"timetracker/src/lib"
	taskModel "timetracker/src/model/task"

	"github.com/gin-gonic/gin"
)

type CompleteTaskBodyParam struct {
	TaskId int `json:"taskId"`
}

type CompleteTaskRes struct {
	TaskId int `json:"taskId"`
}

func completeTask(c *gin.Context) {
	var b CompleteTaskBodyParam

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(400, "invalid body")
		return
	}

	err := taskModel.CompleteTask(b.TaskId)
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
