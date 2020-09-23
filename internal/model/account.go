package model

import (
	"github.com/murilosrg/financial-api/internal/database"
)

type Account struct {
	GORMBase
	Document string `json:"document_number" gorm:"type:varchar(11);unique_index;not null"`
}

func (Account) TableName() string {
	return "account"
}

func (a *Account) Create() error {
	if err := database.DB.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a *Account) Find(id int) error {
	if err := database.DB.First(&a, id).Error; err != nil {
		return err
	}

	return nil
}
