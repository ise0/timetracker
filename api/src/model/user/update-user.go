package userModel

import (
	"context"
	"fmt"
	"strings"
	"timetracker/src/db"
	"timetracker/src/lib"
)

type UpdateUserUserParam struct {
	UserId               int
	Name                 string
	Surname              string
	Patronymic           string
	Address              string
	PassportNumber       string
	NameUpdate           bool
	SurnameUpdate        bool
	PatronymicUpdate     bool
	AddressUpdate        bool
	PassportNumberUpdate bool
}

func UpdateUser(user UpdateUserUserParam) error {
	var update []string
	var updateValues = []any{user.UserId}
	if user.NameUpdate {
		if user.Name == "" {
			return lib.ApiError400("provide valid name")
		}
		updateValues = append(updateValues, user.Name)
		update = append(update, fmt.Sprintf("user_name = $%v", len(updateValues)))
	}
	if user.SurnameUpdate {
		if user.Surname == "" {
			return lib.ApiError400("provide valid surname")
		}
		updateValues = append(updateValues, user.Surname)
		update = append(update, fmt.Sprintf("surname = $%v", len(updateValues)))
	}
	if user.PatronymicUpdate {
		updateValues = append(updateValues, user.Patronymic)
		update = append(update, fmt.Sprintf("patronymic = $%v", len(updateValues)))
	}
	if user.AddressUpdate {
		if user.Address == "" {
			return lib.ApiError400("provide valid address")
		}
		updateValues = append(updateValues, user.Address)
		update = append(update, fmt.Sprintf("address = $%v", len(updateValues)))
	}
	if user.PassportNumberUpdate {
		if user.Address == "" {
			return lib.ApiError400("provide valid passportNumber")
		}
		updateValues = append(updateValues, user.PassportNumber)
		update = append(update, fmt.Sprintf("passport_number = $%v", len(updateValues)))
	}
	if len(update) == 0 {
		return nil
	}
	if _, err := db.DB.Exec(context.TODO(), `
		update users set
		`+strings.Join(update, ", ")+`
		where user_id = $1 and not deleted
	`, updateValues...); err != nil {
		lib.Logger.Errorw("Failed to execute sql query")
		return err
	}

	return nil
}
