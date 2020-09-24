package accounts

import "github.com/murilosrg/financial-api/internal/model"

type Account struct {
	model.GORMBase
	Document string `json:"document_number" gorm:"type:varchar(11);unique_index;not null"`
}

func (Account) TableName() string {
	return "account"
}
