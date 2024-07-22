package app

import (
	"context"
	"fmt"
	"os"
	"time"
	"timetracker/src/api"
	"timetracker/src/db"
	"timetracker/src/lib"
)

func Start() {
	err := lib.Retry(context.TODO(), func() error {
		return db.Connect()
	}, lib.RetryOptions{Retries: -1, Delay: time.Second * 2})
	if err != nil {
		return
	}

	port := os.Getenv("PORT")

	lib.Logger.Info(fmt.Sprintf("Server listening and serving HTTP on port: %v", port))

	err = api.Engine.Run(":" + port)

	if err != nil {
		lib.Logger.Errorw("HTTP server stopped", "error", err)
		return
	}
}
