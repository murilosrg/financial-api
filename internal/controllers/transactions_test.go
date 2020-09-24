package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/murilosrg/financial-api/internal/controllers"
	"github.com/murilosrg/financial-api/internal/model"
	"github.com/murilosrg/financial-api/internal/model/mocks"
	"github.com/murilosrg/financial-api/internal/model/transactions"
	"github.com/shopspring/decimal"
)

func TestPostTransaction(t *testing.T) {
	cases := []struct {
		name        string
		transaction *transactions.Transaction
		err         error
		expectedErr *gin.H
		statusCode  int
	}{
		{
			name: "Success",
			transaction: &transactions.Transaction{
				GORMBase: model.GORMBase{
					ID:        1,
					CreatedAt: time.Now().UTC(),
				},
				Amount:          decimal.NewFromFloat32(10.5),
				AccountID:       1,
				OperationTypeID: 1,
			},
			err:         nil,
			expectedErr: nil,
			statusCode:  http.StatusCreated,
		},
		{
			name:        "Failed - BAD REQUEST",
			transaction: nil,
			err:         errors.New("invalid transaction"),
			expectedErr: &gin.H{"error": "invalid transaction"},
			statusCode:  http.StatusBadRequest,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockITransactionService(ctrl)
			service.EXPECT().Create(gomock.Any()).Return(tt.transaction, tt.err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			bodyBytes, _ := json.Marshal(tt.transaction)
			request := &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(bodyBytes)),
			}

			c.Request = request

			controller := controllers.NewTransactionController(service)
			controller.Post(c)

			response, actual := extractResponseFromResponseWriter(w)

			var actualTransaction *transactions.Transaction
			var actualErr *gin.H

			if tt.statusCode == http.StatusCreated {
				_ = json.Unmarshal([]byte(actual), &actualTransaction)
			} else {
				_ = json.Unmarshal([]byte(actual), &actualErr)
			}

			assert.Equal(t, tt.statusCode, response.StatusCode)
			assert.Equal(t, tt.transaction, actualTransaction)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}
