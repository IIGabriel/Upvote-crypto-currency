package main

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
)

func main() {
	db := config.OpenConnection()
	config.Migrations(db)
	currency := models.Currency{Name: "Bitcoin"}
	currency.FindBy(db)
	services.GetPrice(&currency)
}
