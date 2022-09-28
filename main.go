package main

import (
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
)

func main() {
	db := config.OpenConnection()
	config.Migrations(db)
	test := models.Currency{Name: "klever", Symbol: "klv"}
	if err := test.Create(); err != nil {
		fmt.Println(err)
	}
	services.GetPrice(&test)
}
