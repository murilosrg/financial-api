package accounts_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/murilosrg/financial-api/internal/model"
	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/murilosrg/financial-api/internal/model/mocks"
	"github.com/stretchr/testify/assert"
)

func TestServiceFind(t *testing.T) {
	cases := []struct {
		name    string
		account *accounts.Account
		err     error
	}{
		{
			name:    "Sucess",
			account: &accounts.Account{GORMBase: model.GORMBase{ID: 1, CreatedAt: time.Now()}, Document: "12345678901"},
			err:     nil,
		},
		{
			name:    "Failed",
			account: nil,
			err:     errors.New("account not found"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockIAccountRepository(ctrl)
			repo.EXPECT().Find(gomock.Any()).Return(tt.account, tt.err)

			service := accounts.NewAccountService(repo)
			actualAccount, actualErr := service.Find(1)

			assert.True(t, assertAccount(actualAccount, tt.account))
			assert.Equal(t, actualErr, tt.err)
		})
	}
}

func TestServiceCreate(t *testing.T) {
	cases := []struct {
		name            string
		account         *accounts.Account
		repoErr         error
		expectedAccount *accounts.Account
		expectedErr     error
	}{
		{
			name:            "Sucess",
			account:         &accounts.Account{GORMBase: model.GORMBase{ID: 1, CreatedAt: time.Now()}, Document: "12345678901"},
			repoErr:         nil,
			expectedAccount: &accounts.Account{GORMBase: model.GORMBase{ID: 1, CreatedAt: time.Now()}, Document: "12345678901"},
			expectedErr:     nil,
		},
		{
			name:            "Failed - Error creating",
			account:         &accounts.Account{GORMBase: model.GORMBase{ID: 2, CreatedAt: time.Now()}, Document: "12345678901"},
			repoErr:         errors.New("create failed"),
			expectedAccount: nil,
			expectedErr:     errors.New("create failed"),
		},
		{
			name:            "Failed - Error document invalid",
			account:         &accounts.Account{GORMBase: model.GORMBase{ID: 3, CreatedAt: time.Now()}, Document: "invalid"},
			repoErr:         nil,
			expectedAccount: nil,
			expectedErr:     errors.New("invalid document"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockIAccountRepository(ctrl)

			if tt.account.ID != 3 {
				repo.EXPECT().Create(gomock.Any()).Return(tt.repoErr)
			}

			service := accounts.NewAccountService(repo)
			actualAccount, actualErr := service.Create(tt.account)

			assert.True(t, assertAccount(actualAccount, tt.expectedAccount))
			assert.Equal(t, actualErr, tt.expectedErr)
		})
	}
}

func assertAccount(actual, expected *accounts.Account) bool {
	if actual == nil && expected == nil {
		return true
	}

	if actual != nil && expected != nil {
		return actual.ID == expected.ID && actual.Document == expected.Document
	}

	return false
}
