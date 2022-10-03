package models

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Vote struct {
	Type       string   `json:"type" bson:"type" gorm:"type:varchar(4)"`
	CurrencyId uint     `json:"currency_id" bson:"currency_id"`
	Currency   Currency `gorm:"foreignKey:CurrencyId"`
}

func (c *Currency) CreateUpVote(db *gorm.DB) error {
	vote := Vote{Type: "UP", CurrencyId: c.Id}

	if err := db.Table("votes").Create(&vote).Error; err != nil {
		zap.L().Warn("Error Votes - CreateUpVote():", zap.Error(err))
		return err
	}
	return nil
}

func (c *Currency) CreateDownVote(db *gorm.DB) error {
	vote := Vote{Type: "DOWN", CurrencyId: c.Id}

	if err := db.Table("votes").Create(&vote).Error; err != nil {
		zap.L().Warn("Error Votes - CreateDownVote():", zap.Error(err))
		return err
	}
	return nil
}

func (c *Currency) FindVotes(db *gorm.DB) error {
	if err := db.Raw("select (select count(type) from votes where type = 'UP' and currency_id = ?) UP, (select count(type) from votes where type = 'DOWN' and currency_id = ?) DOWN", c.Id, c.Id).Scan(&c.Votes).Error; err != nil {
		zap.L().Warn("Error Votes - FindVotes():", zap.Error(err))
		return err
	}
	return nil
}
