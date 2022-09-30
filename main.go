package main

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/routes"
	"log"
)

func main() {
	//db := config.OpenConnection()
	//defer config.CloseConnection(db)
	//config.Migrations(db)

	app := routes.AllRoutes()

	if err := app.Listen(":7777"); err != nil {
		log.Fatal(err)
	}

}
