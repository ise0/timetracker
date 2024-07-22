package taskModel

import (
	"context"
	"timetracker/src/db"
	"timetracker/src/lib"
)

type AddTaskRes struct {
	TaskId int
}

func AddTask(taskName string, executorId int) (AddTaskRes, error) {
	var res AddTaskRes

	if taskName == "" {
		return res, lib.ApiError400("provide valid task name")
	}

	if err := db.DB.QueryRow(context.TODO(), `
		insert into tasks(task_name, executor, created_at) values 
		($1, $2, now())
		returning task_id
	`, taskName, executorId).Scan(&res.TaskId); err != nil {
		lib.Logger.Errorw("Failed to execute sql query", err)
		return res, err
	}

	return res, nil
}
