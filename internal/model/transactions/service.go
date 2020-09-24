package transactions

import (
	"errors"

	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/murilosrg/financial-api/internal/model/operations"
)

type ITransactionService interface {
	Create(transaction *Transaction) (*Transaction, error)
}

type TransactionService struct {
	account         accounts.IAccountService
	transactionRepo ITransactionRepository
	operationRepo   operations.IOperationTypeRepository
}

func NewTransactionService(
	account accounts.IAccountService,
	transactionRepo ITransactionRepository,
	operationRepo operations.IOperationTypeRepository) ITransactionService {
	return &TransactionService{
		account:         account,
		transactionRepo: transactionRepo,
		operationRepo:   operationRepo,
	}
}

func (t *TransactionService) Create(transaction *Transaction) (*Transaction, error) {
	_, err := t.account.Find(transaction.AccountID)

	if err != nil {
		return nil, errors.New("account not found")
	}

	operation, err := t.operationRepo.Find(transaction.OperationTypeID)
	if err != nil {
		return nil, errors.New("operation invalid")
	}

	if operation.IsDebit && transaction.Amount.IsPositive() {
		transaction.Amount = transaction.Amount.Neg()
	}

	if err := t.transactionRepo.Create(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}
