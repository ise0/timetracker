package userModel

import (
	"context"
	"time"
	"timetracker/src/db"
	"timetracker/src/lib"
)

type GetUserStatRes struct {
	TaskId    int
	TaskName  string
	SpentTime time.Duration
	Completed bool
}

func GetUserStat(userId int) ([]GetUserStatRes, error) {
	var res []GetUserStatRes

	rows, err := db.DB.Query(context.TODO(), `
		select 
			task_id,
			task_name,
			case when completed_at is null then
				now() - created_at
			else 
				completed_at - created_at
			end as spent_time,
			completed_at is not null as completed 
		from tasks
		join users on executor = user_id
		where executor = $1 and not deleted 
		order by spent_time desc`, userId)

	if err != nil {
		lib.Logger.Errorw("Failed to execute sql query", err)
		return nil, lib.ApiError500("something went wrong")
	}

	defer rows.Close()

	for rows.Next() {
		var resItem GetUserStatRes
		err = rows.Scan(&resItem.TaskId, &resItem.TaskName, &resItem.SpentTime, &resItem.Completed)

		if err != nil {
			lib.Logger.Errorw("Failed to execute sql query", err)
			return nil, lib.ApiError500("")
		}

		res = append(res, resItem)
	}

	if rows.Err() != nil {
		lib.Logger.Errorw("Failed to execute sql query", err)
		return nil, lib.ApiError500("")
	}

	return res, nil
}
