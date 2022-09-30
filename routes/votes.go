package routes

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/gofiber/fiber/v2"
)

func CreateUpVote(c *fiber.Ctx) error {
	coin, err := ValidCurrency(c)
	if err != nil {
		return err // Status defined in method
	}

	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err := coin.CreateUpVote(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error")
	}

	return c.Status(fiber.StatusCreated).JSON("Upvote registered")
}

func CreateDownVote(c *fiber.Ctx) error {
	coin, err := ValidCurrency(c)
	if err != nil {
		return err // Status defined in method
	}

	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err = coin.CreateDownVote(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error")
	}

	return c.Status(fiber.StatusCreated).JSON("The vote has been registered")
}
