package main

import (
	"jummechu/api/routes"
	"jummechu/lib"
)

func main() {
	server := lib.NewServer()

	menuRoutes := routes.NewMenuRoutes(server)
	menuRoutes.Setup()

	server.Run()
}
