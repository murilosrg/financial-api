package operations_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/murilosrg/financial-api/internal/model/operations"
	"github.com/stretchr/testify/assert"
)

func TestRepoFind(t *testing.T) {
	cases := []struct {
		name      string
		operation *operations.OperationType
		err       error
	}{
		{
			name:      "Success",
			operation: &operations.OperationType{ID: 1, Description: "Saque"},
			err:       nil,
		},
		{
			name:      "Failed",
			operation: nil,
			err:       errors.New("record not found"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			DB, _ := gorm.Open("sqlite3", db)

			var rows *sqlmock.Rows

			if tt.err == nil {
				rows = sqlmock.
					NewRows([]string{"id", "description"}).
					AddRow(tt.operation.ID, tt.operation.Description)
			} else {
				rows = sqlmock.NewRows([]string{"id", "description"})
			}

			escapedStmt := regexp.QuoteMeta(`SELECT`)
			mock.ExpectQuery(escapedStmt).
				WillReturnRows(rows)

			repo := operations.NewOperationTypeRepository(DB)
			actualOperation, actualErr := repo.Find(1)

			assert.Equal(t, actualOperation, tt.operation)
			assert.Equal(t, actualErr, tt.err)
		})
	}
}

func TestRepoCreate(t *testing.T) {
	cases := []struct {
		name      string
		operation *operations.OperationType
		err       error
	}{
		{
			name:      "Success",
			operation: &operations.OperationType{ID: 1, Description: "Saque"},
			err:       nil,
		},
		{
			name:      "Failed",
			operation: &operations.OperationType{ID: 1, Description: "Saque"},
			err:       errors.New("error creating"),
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

			repo := operations.NewOperationTypeRepository(DB)
			actualErr := repo.Create(tt.operation)

			assert.Equal(t, actualErr, tt.err)
		})
	}
}
