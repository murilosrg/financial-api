package transactions_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	"github.com/murilosrg/financial-api/internal/model"
	"github.com/murilosrg/financial-api/internal/model/transactions"
	"github.com/shopspring/decimal"
)

func TestRepoCreate(t *testing.T) {
	cases := []struct {
		name        string
		transaction *transactions.Transaction
		err         error
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
			err: nil,
		},
		{
			name:        "failed - error creating",
			transaction: &transactions.Transaction{},
			err:         errors.New("record invalid"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			DB, _ := gorm.Open("sqlite3", db)

			escapedStmt := regexp.QuoteMeta(`INSERT`)
			mock.ExpectBegin()

			if tt.err == nil {
				mock.ExpectExec(escapedStmt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec(escapedStmt).WillReturnError(tt.err)
			}

			mock.ExpectCommit()

			repo := transactions.NewTransactionRepository(DB)
			actualErr := repo.Create(tt.transaction)

			assert.Equal(t, actualErr, tt.err)
		})
	}
}
