package userModel

import (
	"context"
	"fmt"
	"strings"
	"timetracker/src/db"
	"timetracker/src/lib"
)

type GetUsersPaginationParam struct {
	Limit  int
	Offset int
}

type GetUsersFiltersParam struct {
	UserId         []int
	Name           string
	Surname        string
	Patronymic     string
	PassportNumber string
	Address        string
}

type GetUsersRes struct {
	UserId         int
	Name           string
	Surname        string
	Patronymic     string
	PassportNumber string
	Address        string
}

func GetUsers(pagination GetUsersPaginationParam,
	filters GetUsersFiltersParam) ([]GetUsersRes, error) {

	var f []string
	var args = []any{pagination.Limit, pagination.Offset}
	if filters.Name != "" {
		args = append(args, "%"+filters.Name+"%")
		f = append(f, fmt.Sprintf("user_name like $%v", len(args)))
	}
	if filters.Surname != "" {
		args = append(args, "%"+filters.Surname+"%")
		f = append(f, fmt.Sprintf("surname like $%v", len(args)))
	}
	if filters.Patronymic != "" {
		args = append(args, "%"+filters.Patronymic+"%")
		f = append(f, fmt.Sprintf("patronymic like $%v", len(args)))
	}
	if filters.PassportNumber != "" {
		args = append(args, "%"+filters.PassportNumber+"%")
		f = append(f, fmt.Sprintf("passport_number like $%v", len(args)))
	}
	if filters.Address != "" {
		args = append(args, "%"+filters.Address+"%")
		f = append(f, fmt.Sprintf("address like $%v", len(args)))
	}

	if len(filters.UserId) > 0 {
		args = append(args, filters.UserId)
		f = append(f, fmt.Sprintf("user_id = any ($%v)", len(args)))
	}

	var filtersSqlStr string
	if len(f) > 0 {
		filtersSqlStr += " and " + strings.Join(f, " and ")
	}

	var res []GetUsersRes
	rows, err := db.DB.Query(context.TODO(), `
		select 
			user_id,
			user_name,
			surname,
			patronymic,
			address,
			passport_number
		from users
		where not deleted
		`+filtersSqlStr+`
		limit $1 offset $2`, args...)

	if err != nil {
		lib.Logger.Errorw("Failed to execute sql query", err)
		return nil, lib.ApiError500("")
	}

	defer rows.Close()
	for rows.Next() {
		var resItem GetUsersRes
		err = rows.Scan(&resItem.UserId, &resItem.Name, &resItem.Surname,
			&resItem.Patronymic, &resItem.Address, &resItem.PassportNumber)

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
