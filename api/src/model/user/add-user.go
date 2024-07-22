package userModel

import (
	"context"
	"timetracker/src/db"
	"timetracker/src/lib"

	"github.com/jackc/pgx/v5"
)

type AddUserParam struct {
	Name           string
	Surname        string
	Patronymic     string
	Address        string
	PassportNumber string
}

type AddUserRes struct {
	UserId int
}

func AddUser(user AddUserParam) (AddUserRes, error) {
	var res AddUserRes

	sqlFailLog := func(err error) (AddUserRes, error) {
		lib.Logger.Errorw("Failed to execute sql query", "error", err)
		return res, err
	}

	if user.Name == "" {
		return res, lib.ApiError400("provide valid name")
	} else if user.Surname == "" {
		return res, lib.ApiError400("provide valid surname")
	} else if user.Address == "" {
		return res, lib.ApiError400("provide valid address")
	}

	tx, err := db.DB.BeginTx(context.TODO(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	defer tx.Rollback(context.TODO())

	if err != nil {
		return sqlFailLog(err)
	}

	var uniquePassport bool
	if err := tx.QueryRow(context.TODO(), `
		select not exists (select user_id from users where passport_number = $1)
	`, user.PassportNumber).
		Scan(&uniquePassport); err != nil {
		return sqlFailLog(err)
	}

	if !uniquePassport {
		return res, lib.ApiError400("provided passport already exists")
	}

	if err := tx.QueryRow(context.TODO(), `
		insert into users(user_name, surname, patronymic, address, passport_number) 
		values ($1, $2, $3, $4, $5)
		returning user_id
	`, user.Name, user.Surname, user.Patronymic, user.Address, user.PassportNumber).
		Scan(&res.UserId); err != nil {
		return sqlFailLog(err)
	}

	if err = tx.Commit(context.TODO()); err != nil {
		return sqlFailLog(err)
	}

	return res, nil
}
