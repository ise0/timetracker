package db

import (
	"context"
	"os"
	"timetracker/src/lib"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *pgx.Conn

func Connect() error {
	d, err := pgx.Connect(context.TODO(), os.Getenv("DB_CONNECTION"))

	if err != nil {
		lib.Logger.Errorw("Failed to connect to db", "error", err)
		return err
	}
	DB = d

	lib.Logger.Info("Connected to db successfully")

	var initRequired bool

	if err := DB.QueryRow(context.TODO(), `
		select count(*) = 0
		from information_schema.tables
		where table_schema = 'public'
	`).Scan(&initRequired); err != nil {
		DB.Close(context.TODO())
		lib.Logger.Errorw("Failed to execute sql query", "error", err)
		return err
	}

	if initRequired {
		if _, err := DB.Exec(context.TODO(), dbSchema); err != nil {
			d.Close(context.TODO())
			lib.Logger.Errorw("Failed to execute sql query", "error", err)
			return err
		}
	}

	return nil
}
