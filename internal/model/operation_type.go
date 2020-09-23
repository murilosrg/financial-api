package model

import (
	"github.com/murilosrg/financial-api/internal/database"
)

type OperationType struct {
	ID          int    `gorm:"primaryKey"`
	Description string `json:"description" gorm:"not null;type:varchar(50)"`
	IsDebit     bool   `json:"is_debit" gorm:"not null;column:debit"`
}

func (OperationType) TableName() string {
	return "operation_type"
}

func (o *OperationType) Create() error {
	if err := database.DB.Create(&o).Error; err != nil {
		return err
	}

	return nil
}

func (o *OperationType) Find(id int) error {
	if err := database.DB.First(&o, id).Error; err != nil {
		return err
	}

	return nil
}
