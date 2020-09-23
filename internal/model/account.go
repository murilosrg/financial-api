package model

import (
	"github.com/murilosrg/financial-api/internal/database"
)

type Account struct {
	GORMBase
	Document string `json:"document" gorm:"type:varchar(11);unique_index;not null"`
}

func (Account) TableName() string {
	return "account"
}

func (a *Account) Create() (account *Account, err error) {
	if err := database.DB.Create(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}
