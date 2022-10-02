package routes

import (
	"encoding/json"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/IIGabriel/Upvote-crypto-currency.git/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetCurrency(c *fiber.Ctx) error {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	coin, err := models.ValidCurrency(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params")
	}

	if c.Query("omit_price") != "true" {
		if err = services.GetPrice(&coin, c); err != nil {
			zap.L().Info("Error Currency - GetPrice():", zap.Error(err))
		}
	}

	if err = coin.FindVotes(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error")
	}

	return c.Status(fiber.StatusOK).JSON(coin)

}

func CreateCurrency(c *fiber.Ctx) error {
	if !config.ValidatorSessionToken(c) {
		return c.Status(fiber.StatusUnauthorized).JSON("Permission denied")
	}

	var coin models.Currency

	if err := json.Unmarshal(c.Body(), &coin); err != nil {
		zap.L().Info("Error JSON Unmarshal - CreateCurrency():", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Error")
	}

	if coin.CoinId == "" || coin.Name == "" || coin.Symbol == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Inform all fields")
	}

	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err := coin.CreateIfNotExist(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Error")
	}

	if coin.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("Currency already exist")
	}

	return c.Status(fiber.StatusCreated).JSON("Currency added")

}

func DeleteCurrency(c *fiber.Ctx) error {
	if !config.ValidatorSessionToken(c) {
		return c.Status(fiber.StatusUnauthorized).JSON("Permission denied")
	}
	db := config.OpenConnection()
	defer config.CloseConnection(db)
	coin, err := models.ValidCurrency(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Params")
	}

	if err = coin.Delete(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Internal error")
	}

	return c.Status(fiber.StatusOK).JSON("Currency deleted")

}

func EditCurrency(c *fiber.Ctx) error {
	if !config.ValidatorSessionToken(c) {
		return c.Status(fiber.StatusUnauthorized).JSON("Permission denied")
	}
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	coin, err := models.ValidCurrency(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Params")
	}
	if err = json.Unmarshal(c.Body(), &coin); err != nil {
		zap.L().Info("Error JSON Unmarshal - CreateCurrency():", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Error")
	}

	if err = coin.Update(db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not update")
	}

	return c.Status(fiber.StatusOK).JSON("Currency updated")
}
