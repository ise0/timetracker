package taskRouter

import (
	"timetracker/src/lib"
	taskModel "timetracker/src/model/task"

	"github.com/gin-gonic/gin"
)

type addTaskBodyParam struct {
	TaskName string `json:"taskName"`
	Executor int    `json:"executor"`
}

type AddTaskRes struct {
	TaskId int `json:"task_id"`
}

func addTask(c *gin.Context) {
	var b addTaskBodyParam

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(400, "invalid body")
		return
	}

	r, err := taskModel.AddTask(b.TaskName, b.Executor)
	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	c.JSON(200, AddTaskRes{r.TaskId})
}
