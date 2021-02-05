package main

import (
	"group-management-api/app"
)

func main() {
	// Init the application.
	con := app.InitApp()
	// Start the application.
	con.StartApp()
	// Defer the shut down off all the queued things.
	defer con.ShutdownAll()
}