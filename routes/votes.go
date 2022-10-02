package routes

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUpVote(c *fiber.Ctx) error {
	db := config.OpenConnection()
	defer config.CloseConnection(db)
	coin, err := models.ValidCurrency(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params")
	}

	if err = coin.CreateUpVote(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Internal error")
	}

	return c.Status(fiber.StatusCreated).JSON("Upvote registered")
}

func CreateDownVote(c *fiber.Ctx) error {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	coin, err := models.ValidCurrency(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params")
	}

	if err = coin.CreateDownVote(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Internal error")
	}

	return c.Status(fiber.StatusCreated).JSON("Downvote registered")
}
