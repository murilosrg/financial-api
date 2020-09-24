package transactions

import (
	"github.com/murilosrg/financial-api/internal/model"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	model.GORMBase
	Amount          decimal.Decimal `json:"amount" gorm:"type:decimal(20,8);not null"`
	AccountID       int             `json:"account_id" gorm:"not null;index"`
	OperationTypeID int             `json:"operation_type_id" gorm:"not null;index"`
}

func (Transaction) TableName() string {
	return "transaction"
}
