package model

import (
	"github.com/murilosrg/financial-api/internal/database"
)

type OperationType struct {
	ID          int    `gorm:"primaryKey"`
	Description string `json:"description" gorm:"not null;type:varchar(50)"`
	IsDebit     bool   `json:"is_debit" gorm:"not null;column:debit;type:bit"`
}

func (OperationType) TableName() string {
	return "operation_type"
}

func (o *OperationType) Create() (operation *OperationType, err error) {
	if err = database.DB.Create(&o).Error; err != nil {
		return nil, err
	}

	return operation, nil
}

func (o *OperationType) Find() (operation *OperationType, err error) {
	if err = database.DB.Where(&o).First(&operation).Error; err != nil {
		return nil, err
	}

	return operation, nil
}
