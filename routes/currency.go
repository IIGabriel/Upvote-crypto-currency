package routes

import (
	"errors"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetCurrency(c *fiber.Ctx) error {
	coin, err := ValidCurrency(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params")
	}

	if err = services.GetPrice(&coin); err != nil {
		zap.L().Info("Error Currency - GetPrice():", zap.Error(err))
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
		return coin, errors.New("Invalid params")
	}

	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err := coin.FindBy(db); err != nil {
		return coin, err
	}

	if coin.Id == 0 {
		return coin, errors.New("Invalid params")
	}
	return coin, nil
}
