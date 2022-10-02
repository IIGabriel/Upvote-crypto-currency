package main

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/routes"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
	"go.uber.org/zap"
)

func main() {
	config.InitLogger()

	///////////////// Migrations /////////////////

	config.Migrations()
	services.GetAllCoins()

	///////////////// Listen /////////////////

	app := routes.AllRoutes()

	if err := app.Listen(":" + config.GetEnv("listen_port")); err != nil {
		zap.L().Panic("Error listening port", zap.Error(err))
	}

}
