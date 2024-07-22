package main

import (
	"timetracker/src/app"
	"timetracker/src/lib"
)

func main() {
	lib.CreateLogger()
	defer lib.Logger.Sync()

	app.Start()
}
