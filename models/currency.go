package models

import (
	"errors"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type Currency struct {
	Id     uint    `gorm:"primary_key" json:"-" bson:"-"`
	CoinId string  `json:"id" bson:"id" gorm:"type:varchar(100)"`
	Name   string  `json:"name" bson:"name" gorm:"type:varchar(100)"`
	Symbol string  `json:"symbol" bson:"symbol" gorm:"type:varchar(100)"`
	Votes  Votes   `json:"votes" bson:"votes" gorm:"-"`
	Prices []Price `json:"prices" bson:"prices" gorm:"-"`
}

type Price struct {
	Date  string  `json:"date" bson:"date"`
	Price float64 `json:"price" bson:"price"`
}

type Votes struct {
	Up   int `json:"up" bson:"up"`
	Down int `json:"down"bson:"down""`
}

func (c *Currency) CreateIfNotExist(db *gorm.DB) error {
	var id int
	c.Name = strings.ToUpper(c.Name)
	db.Table("currencies").Select("id").Where("name = ?", c.Name).Scan(&id)
	if id != 0 {
		return nil
	}

	if err := db.Table("currencies").Create(&c).Where(c).Error; err != nil {
		zap.L().Warn("Error Currency - Create():", zap.Error(err))
		return err
	}
	return nil
}

func (c *Currency) FindBy(db *gorm.DB) error {
	c.Name = strings.ToUpper(c.Name)
	if err := db.Table("currencies").Where(c).Find(&c).Error; err != nil {
		zap.L().Warn("Error Currency - FindBy():", zap.Error(err))
		return err
	}
	return nil
}

func (c *Currency) Delete(db *gorm.DB) error {
	c.Name = strings.ToUpper(c.Name)

	if err := db.Table("currencies").Delete(&c).Error; err != nil {
		zap.L().Warn("Error Currency - Delete():", zap.Error(err))
		return err
	}

	return nil
}

func ValidCurrency(c *fiber.Ctx) (Currency, error) {
	var coin Currency
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
