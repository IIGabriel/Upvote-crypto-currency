package main

import (
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
)

func main() {
	db := config.OpenConnection()
	defer config.CloseConnection(db)
	//config.Migrations(db)
	currency := models.Currency{Name: "Bitcoin"}
	if err := currency.FindByName(db); err != nil {
		fmt.Println("Erro")
	}
	services.GetPrice(&currency)
	fmt.Println(currency)

	currency.CreateDownVote(db)
	currency.CreateUpVote(db)
	err := currency.FindVotes(db)
	if err != nil {
		fmt.Println("Erro")
	}
	fmt.Println(currency)
}
