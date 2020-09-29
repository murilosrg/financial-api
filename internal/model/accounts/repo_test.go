package accounts_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/murilosrg/financial-api/internal/model"
	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/stretchr/testify/assert"
)

func TestRepoFind(t *testing.T) {
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
			name:    "failed - not found",
			account: nil,
			err:     errors.New("record not found"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			DB, _ := gorm.Open("sqlite3", db)

			var rows *sqlmock.Rows

			if tt.err == nil {
				rows = sqlmock.
					NewRows([]string{"id", "document", "created_at"}).
					AddRow(tt.account.ID, tt.account.Document, tt.account.CreatedAt)
			} else {
				rows = sqlmock.NewRows([]string{"id", "document", "created_at"})
			}

			escapedStmt := regexp.QuoteMeta(`SELECT`)
			mock.ExpectQuery(escapedStmt).
				WillReturnRows(rows)

			repo := accounts.NewAccountRepository(DB)
			actualAccount, actualErr := repo.Find(1)

			assert.Equal(t, actualAccount, tt.account)
			assert.Equal(t, actualErr, tt.err)
		})
	}
}

func TestRepoCreate(t *testing.T) {
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
			name:    "failed - error creating",
			account: &accounts.Account{},
			err:     errors.New("record invalid"),
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

			repo := accounts.NewAccountRepository(DB)
			actualErr := repo.Create(tt.account)

			assert.Equal(t, actualErr, tt.err)
		})
	}
}
