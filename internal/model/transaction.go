package model

import (
	"github.com/murilosrg/financial-api/internal/database"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	GORMBase
	Amount          decimal.Decimal `json:"amount" gorm:"type:decimal(20,8);not null"`
	AccountID       int             `json:"account_id" gorm:"not null;index"`
	OperationTypeID int             `json:"operation_type_id" gorm:"not null;index"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func (t *Transaction) Create() (transaction *Transaction, err error) {
	if err := database.DB.Create(&t).Error; err != nil {
		return nil, err
	}

	return t, err
}
