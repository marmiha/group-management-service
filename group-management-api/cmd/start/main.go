package main

import (
	"group-management-api/app"
	"group-management-api/app/logger"
)

func main() {
	// Init the application.
	con, err := app.InitApp()
	if err != nil {
		logger.Log.WithField("err", err).Fatal("Something went wrong initializing the application.")
	}

	// Start the application.
	if err := con.StartApp(); err != nil {
		logger.Log.WithField("err", err).Fatal("Could not start application.")
	}

	// Defer the shut down off all the queued things.
	defer con.ShutdownAll()
}