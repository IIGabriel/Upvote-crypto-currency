package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Vote struct {
	Id         uint     `json:"id" bson:"id" gorm:"primary_key"`
	Type       string   `json:"type" bson:"type" gorm:"type:varchar(4)"`
	CurrencyId uint     `json:"currency_id" bson:"currency_id"`
	Currency   Currency `json:"currency" bson:"currency" gorm:"type:foreignKey:CurrencyId"`
}

func (c *Currency) CreateUpVote(db *gorm.DB) error {
	vote := Vote{Type: "UP", CurrencyId: c.Id}

	if err := db.Table("votes").Create(&vote).Error; err != nil {
		fmt.Println("erro")
		return err
	}
	return nil
}
func (c *Currency) CreateDownVote(db *gorm.DB) error {
	vote := Vote{Type: "DOWN", CurrencyId: c.Id}

	if err := db.Table("votes").Create(&vote).Error; err != nil {
		fmt.Println("erro")
		return err
	}
	return nil
}
