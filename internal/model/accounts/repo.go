package accounts

import (
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source=repo.go -destination=../mocks/mock_account_repository.go -package=mocks

type IAccountRepository interface {
	Create(account *Account) error
	Find(id int) (*Account, error)
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

func (a *AccountRepository) Find(id int) (*Account, error) {
	var account Account
	if err := a.db.First(&account, id).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
