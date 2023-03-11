package main

import (
	"echo-user-app/db"
	"echo-user-app/routes"
)

func main() {
	db.Init()

	e := routes.InitRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
