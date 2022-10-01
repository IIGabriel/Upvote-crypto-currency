package main

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/routes"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
	"go.uber.org/zap"
)

func main() {
	///////////////// Migrations /////////////////

	db := config.OpenConnection()
	if err := db.AutoMigrate(models.Currency{}, models.Vote{}); err != nil {
		zap.L().Panic("Could not do migrations", zap.Error(err))
	}
	if err := services.GetAllCoins(db); err != nil {
		zap.L().Panic("Could not get currency values", zap.Error(err))
	}
	config.CloseConnection(db)

	///////////////// Listen /////////////////

	app := routes.AllRoutes()

	if err := app.Listen(":" + config.GetEnv("listen_port")); err != nil {
		zap.L().Panic("Error listening port", zap.Error(err))
	}

}
