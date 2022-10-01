package main

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/routes"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
)

func main() {
	///////////////// Migrations /////////////////

	db := config.OpenConnection()
	if err := db.AutoMigrate(models.Currency{}, models.Vote{}); err != nil {
		panic(err)
	}
	if err := services.GetAllCoins(db); err != nil {
		panic(err)
	}
	config.CloseConnection(db)

	///////////////// Listen /////////////////

	app := routes.AllRoutes()

	if err := app.Listen(":7777"); err != nil {
		panic(err)
	}

}
