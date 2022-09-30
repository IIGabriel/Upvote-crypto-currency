package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Currency struct {
	Id     uint    `gorm:"primary_key"`
	CoinId string  `json:"id" bson:"id" gorm:"type:varchar(100)"`
	Name   string  `json:"name" bson:"name" gorm:"type:varchar(100)"`
	Symbol string  `json:"symbol" bson:"symbol" gorm:"type:varchar(8)"`
	Prices []Price `json:"prices" bson:"prices"`
	Votes  Votes   `json:"votes" bson:"votes"`
}

type Price struct {
	Date  string  `json:"date" bson:"date"`
	Price float64 `json:"price" bson:"price"`
}

type Votes struct {
	Up   int `json:"up" bson:"up"`
	Down int `json:"down"bson:"down""`
}

func (c *Currency) Create(db *gorm.DB) error {
	if err := db.Table("currencies").Create(&c).Where(c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *Currency) FindByName(db *gorm.DB) error {
	if err := db.Table("currencies").Where("name = ?", c.Name).Find(&c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
