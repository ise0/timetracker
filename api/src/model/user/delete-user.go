package userModel

import (
	"context"
	"timetracker/src/db"
	"timetracker/src/lib"
)

func DeleteUser(userId int) (err error) {
	if _, err := db.DB.Exec(context.TODO(), `
		with refs as  (
			select executor from tasks where executor = $1
		), deleted as (
			delete from users 
			where user_id = $1 and not exists (select * from refs) and not deleted
			returning user_id
		)
		update users set
		deleted = true
		where user_id = $1 and not exists (select * from deleted)
	`, userId); err != nil {
		lib.Logger.Errorw("Failed to execute sql query", err)
		return err
	}

	return nil
}
