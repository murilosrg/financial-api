package transactions_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/murilosrg/financial-api/internal/model"
	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/murilosrg/financial-api/internal/model/mocks"
	"github.com/murilosrg/financial-api/internal/model/operations"
	"github.com/murilosrg/financial-api/internal/model/transactions"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreate(t *testing.T) {
	cases := []struct {
		name                string
		transaction         *transactions.Transaction
		accountErr          error
		operationRepoErr    error
		transactionRepoErr  error
		expectedTransaction *transactions.Transaction
		expectedErr         error
	}{
		{
			name: "Sucess",
			transaction: &transactions.Transaction{
				GORMBase: model.GORMBase{
					ID:        1,
					CreatedAt: time.Now().UTC(),
				},
				Amount:          decimal.NewFromFloat32(10.5),
				AccountID:       1,
				OperationTypeID: 1,
			},
			accountErr:         nil,
			operationRepoErr:   nil,
			transactionRepoErr: nil,
			expectedTransaction: &transactions.Transaction{
				GORMBase: model.GORMBase{
					ID:        1,
					CreatedAt: time.Now().UTC(),
				},
				Amount:          decimal.NewFromFloat32(10.5),
				AccountID:       1,
				OperationTypeID: 1,
			},
			expectedErr: nil,
		},
		{
			name: "Failed - Error creating",
			transaction: &transactions.Transaction{
				GORMBase: model.GORMBase{
					ID:        2,
					CreatedAt: time.Now().UTC(),
				},
				Amount:          decimal.NewFromFloat32(10.5),
				AccountID:       1,
				OperationTypeID: 1,
			},
			accountErr:          nil,
			operationRepoErr:    nil,
			transactionRepoErr:  errors.New("create failed"),
			expectedTransaction: nil,
			expectedErr:         errors.New("create failed"),
		},
		{
			name: "Failed - Error account invalid",
			transaction: &transactions.Transaction{
				GORMBase: model.GORMBase{
					ID:        3,
					CreatedAt: time.Now().UTC(),
				},
				Amount:          decimal.NewFromFloat32(10.5),
				AccountID:       1,
				OperationTypeID: 1,
			},
			accountErr:          errors.New("account not found"),
			operationRepoErr:    nil,
			transactionRepoErr:  nil,
			expectedTransaction: nil,
			expectedErr:         errors.New("account not found"),
		},
		{
			name: "Failed - Error operation invalid",
			transaction: &transactions.Transaction{
				GORMBase: model.GORMBase{
					ID:        4,
					CreatedAt: time.Now().UTC(),
				},
				Amount:          decimal.NewFromFloat32(10.5),
				AccountID:       1,
				OperationTypeID: 1,
			},
			accountErr:          nil,
			operationRepoErr:    errors.New("operation invalid"),
			transactionRepoErr:  nil,
			expectedTransaction: nil,
			expectedErr:         errors.New("operation invalid"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			transactionRepo := mocks.NewMockITransactionRepository(ctrl)
			operationRepo := mocks.NewMockIOperationTypeRepository(ctrl)
			accountService := mocks.NewMockIAccountService(ctrl)

			switch tt.transaction.ID {
			case 1, 2:
				accountService.EXPECT().Find(gomock.Any()).Return(&accounts.Account{}, tt.accountErr)
				operationRepo.EXPECT().Find(gomock.Any()).Return(&operations.OperationType{}, tt.operationRepoErr)
				transactionRepo.EXPECT().Create(gomock.Any()).Return(tt.transactionRepoErr)
			case 3:
				accountService.EXPECT().Find(gomock.Any()).Return(&accounts.Account{}, tt.accountErr)
			case 4:
				accountService.EXPECT().Find(gomock.Any()).Return(&accounts.Account{}, tt.accountErr)
				operationRepo.EXPECT().Find(gomock.Any()).Return(&operations.OperationType{}, tt.operationRepoErr)
			}

			service := transactions.NewTransactionService(accountService, transactionRepo, operationRepo)
			actualTransaction, actualErr := service.Create(tt.transaction)

			assert.True(t, assertTransction(actualTransaction, tt.expectedTransaction))
			assert.Equal(t, actualErr, tt.expectedErr)
		})
	}
}

func assertTransction(actual, expected *transactions.Transaction) bool {
	if actual == nil && expected == nil {
		return true
	}

	if actual != nil && expected != nil {
		return actual.ID == expected.ID &&
			actual.AccountID == expected.AccountID &&
			actual.OperationTypeID == expected.OperationTypeID &&
			actual.Amount.Equal(expected.Amount)
	}

	return false
}
