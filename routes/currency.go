package routes

import (
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
	"github.com/gofiber/fiber/v2"
)

func GetCurrency(c *fiber.Ctx) error {
	coin, err := ValidCurrency(c)
	if err != nil {
		return err
	}

	if err = services.GetPrice(&coin); err != nil {
		fmt.Println("Erro")
		//	TODO COLOCAR UMA LOG DESCENTE
	}
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err = coin.FindVotes(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error")
	}

	return c.Status(fiber.StatusOK).JSON(coin)

}

func ValidCurrency(c *fiber.Ctx) (models.Currency, error) {
	var coin models.Currency
	coin.Name = c.Params("coin")

	if coin.Name == "" {
		return coin, c.Status(fiber.StatusBadRequest).JSON("Invalid params")
	}

	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err := coin.FindByName(db); err != nil {
		return coin, c.Status(fiber.StatusInternalServerError).JSON("Error")
	}

	if coin.Id == 0 {
		return coin, c.Status(fiber.StatusBadRequest).JSON("Invalid params")
	}
	return coin, nil
}
