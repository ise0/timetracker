package userModel

import (
	"context"
	"timetracker/src/db"
	"timetracker/src/lib"

	"github.com/jackc/pgx/v5"
)

type GetUserByPassportRes struct {
	Name       string
	Surname    string
	Patronymic string
	Address    string
}

func GetUserByPassport(passportNumber string) (GetUserByPassportRes, error) {
	var res GetUserByPassportRes

	if err := db.DB.QueryRow(context.TODO(), `
		select 
			user_name,
			surname,
			patronymic,
			address
		from users
		where passport_number = $1 and not deleted`, passportNumber).
		Scan(&res.Name, &res.Surname, &res.Patronymic, &res.Address); err != nil {
		if err == pgx.ErrNoRows {
			return res, lib.ApiError400("Bad request")
		}
		lib.Logger.Errorw("Failed to execute sql query", err)
		return res, err
	}

	return res, nil
}
