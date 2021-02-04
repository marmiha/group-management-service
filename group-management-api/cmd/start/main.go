package main

import "group-management-api/app"

func main() {
	con := app.InitApp()
	defer con.ShutdownAll()
}