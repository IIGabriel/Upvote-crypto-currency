package models

import (
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"gorm.io/gorm"
)

type Currency struct {
	Id        uint           `json:"id" bson:"id" gorm:"primary_key"`
	Name      string         `json:"name" bson:"name" gorm:"type:varchar(100)"`
	Symbol    string         `json:"symbol" bson:"symbol" gorm:"type:varchar(8)"`
	Prices    [][]float64    `json:"prices" bson:"prices"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

//type CurrencyInterface interface {
//	GetPrices(Name string) []big.Float
//	Create(currency Currency)
//}

func (c *Currency) Create() error {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	if err := db.Create(&c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
