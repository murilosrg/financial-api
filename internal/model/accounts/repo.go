package accounts

import (
	"github.com/jinzhu/gorm"
)

type IAccountRepository interface {
	Create(account *Account) error
	Find(account *Account, id int) error
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (a *AccountRepository) Create(account *Account) error {
	if err := a.db.Create(&account).Error; err != nil {
		return err
	}

	return nil
}

func (a *AccountRepository) Find(account *Account, id int) error {
	if err := a.db.First(&account, id).Error; err != nil {
		return err
	}

	return nil
}
