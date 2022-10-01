package models

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Currency struct {
	Id     uint    `gorm:"primary_key" json:"-" bson:"-"`
	CoinId string  `json:"coin" bson:"coin" gorm:"type:varchar(100)"`
	Name   string  `json:"name" bson:"name" gorm:"type:varchar(100)"`
	Symbol string  `json:"symbol" bson:"symbol" gorm:"type:varchar(100)"`
	Prices []Price `json:"prices" bson:"prices" gorm:"-"`
	Votes  Votes   `json:"votes" bson:"votes" gorm:"-"`
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
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *Currency) FindBy(db *gorm.DB) error {
	c.Name = strings.ToUpper(c.Name)
	if err := db.Table("currencies").Where(c).Find(&c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
