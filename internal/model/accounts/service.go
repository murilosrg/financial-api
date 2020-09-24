package accounts

import (
	"errors"
	"regexp"
)

//go:generate mockgen -source=service.go -destination=../mocks/mock_account_service.go -package=mocks

type IAccountService interface {
	Find(id int) (*Account, error)
	Create(account *Account) (*Account, error)
}

type AccountService struct {
	repo IAccountRepository
}

func NewAccountService(repo IAccountRepository) IAccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) Find(id int) (*Account, error) {
	account, err := a.repo.Find(id)

	if err != nil {
		return nil, errors.New("account not found")
	}

	return account, nil
}

func (a *AccountService) Create(account *Account) (*Account, error) {
	re := regexp.MustCompile("[0-9]+")

	if !re.MatchString(account.Document) {
		return nil, errors.New("invalid document")
	}

	if err := a.repo.Create(account); err != nil {
		return nil, err
	}

	return account, nil
}
