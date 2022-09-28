package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Currency struct {
	Id     uint        `gorm:"primary_key"`
	CoinId string      `json:"id" bson:"id" gorm:"type:varchar(100)"`
	Name   string      `json:"name" bson:"name" gorm:"type:varchar(100)"`
	Symbol string      `json:"symbol" bson:"symbol" gorm:"type:varchar(8)"`
	Prices [][]float64 `json:"prices" bson:"prices"`
}

func (c *Currency) Create(db *gorm.DB) error {
	if err := db.Table("currencies").Create(&c).Where(c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *Currency) FindBy(db *gorm.DB) error {
	if err := db.Table("currencies").Where("id = 2").Find(&c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
