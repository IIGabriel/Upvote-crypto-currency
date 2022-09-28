package models

type Vote struct {
	Id         uint     `json:"id" bson:"id" gorm:"primary_key"`
	Type       string   `json:"type" bson:"type" gorm:"type:varchar(4)"`
	CurrencyId uint     `json:"currency_id" bson:"currency_id"`
	Currency   Currency `json:"currency" bson:"currency" gorm:"type:foreignKey:CurrencyId"`
}

type VoteInterface interface {
	CreateUpVote(Name int) error
	CreateDownVote(Name int) error
}
