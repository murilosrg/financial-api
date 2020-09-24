package transactions

import (
	"github.com/jinzhu/gorm"
)

type ITransactionRepository interface {
	Create(transaction *Transaction) error
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) Create(transaction *Transaction) error {
	if err := t.db.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
