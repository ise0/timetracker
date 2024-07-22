package taskModel

import (
	"context"
	"timetracker/src/db"
	"timetracker/src/lib"
)

func CompleteTask(taskId int) error {

	if _, err := db.DB.Exec(context.TODO(), `
		update tasks set 
		completed_at = now()
		where task_id = $1 and completed_at is null
		returning task_id
	`, taskId); err != nil {
		lib.Logger.Errorw("Failed to execute sql query", err)
		return err
	}

	return nil
}
