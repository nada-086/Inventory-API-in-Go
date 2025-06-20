package main

import (
	"log"
)

func main() {
	app := App{}
	if err := app.Initialize(DBUser, DBPassword, DBName); err != nil {
		log.Fatal(err)
	}
	app.handleRoutes()
	app.Run("localhost:10000")
}