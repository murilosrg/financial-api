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
	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/murilosrg/financial-api/internal/model/mocks"
)

func TestPostAccount(t *testing.T) {
	cases := []struct {
		name        string
		account     *accounts.Account
		err         error
		expectedErr *gin.H
		statusCode  int
	}{
		{
			name: "Success",
			account: &accounts.Account{
				GORMBase: model.GORMBase{
					ID:        1,
					CreatedAt: time.Now().UTC(),
				},
				Document: "1234567890",
			},
			err:         nil,
			expectedErr: nil,
			statusCode:  http.StatusCreated,
		},
		{
			name:        "Failed - BAD REQUEST",
			account:     nil,
			err:         errors.New("invalid document"),
			expectedErr: &gin.H{"error": "invalid document"},
			statusCode:  http.StatusBadRequest,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockIAccountService(ctrl)
			service.EXPECT().Create(gomock.Any()).Return(tt.account, tt.err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			bodyBytes, _ := json.Marshal(tt.account)
			request := &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(bodyBytes)),
			}

			c.Request = request

			controller := controllers.NewAccountController(service)
			controller.Post(c)

			response, actual := extractResponseFromResponseWriter(w)

			var actualAccount *accounts.Account
			var actualErr *gin.H

			if tt.statusCode == http.StatusCreated {
				_ = json.Unmarshal([]byte(actual), &actualAccount)
			} else {
				_ = json.Unmarshal([]byte(actual), &actualErr)
			}

			assert.Equal(t, tt.statusCode, response.StatusCode)
			assert.Equal(t, tt.account, actualAccount)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}

func TestFindAccount(t *testing.T) {
	cases := []struct {
		name        string
		account     *accounts.Account
		err         error
		expectedErr *gin.H
		params      gin.Params
		statusCode  int
	}{
		{
			name: "Success",
			account: &accounts.Account{
				GORMBase: model.GORMBase{
					ID:        1,
					CreatedAt: time.Now().UTC(),
				},
				Document: "1234567890",
			},
			err:         nil,
			expectedErr: nil,
			params:      gin.Params{{Key: "accountId", Value: "1"}},
			statusCode:  http.StatusOK,
		},
		{
			name:        "Failed - NOT FOUND",
			account:     nil,
			err:         errors.New("account not found"),
			expectedErr: &gin.H{"error": "account not found"},
			statusCode:  http.StatusNotFound,
			params:      gin.Params{{Key: "accountId", Value: "2"}},
		},
		{
			name:        "Failed - INTERNAL ERROR",
			account:     nil,
			err:         errors.New("invalid"),
			expectedErr: &gin.H{"error": "strconv.Atoi: parsing \"invalid\": invalid syntax"},
			statusCode:  http.StatusInternalServerError,
			params:      gin.Params{{Key: "accountId", Value: "invalid"}},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockIAccountService(ctrl)

			if tt.statusCode != http.StatusInternalServerError {
				service.EXPECT().Find(gomock.Any()).Return(tt.account, tt.err)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = tt.params

			controller := controllers.NewAccountController(service)
			controller.Find(c)

			response, actual := extractResponseFromResponseWriter(w)

			var actualAccount *accounts.Account
			var actualErr *gin.H

			if tt.statusCode == http.StatusOK {
				_ = json.Unmarshal([]byte(actual), &actualAccount)
			} else {
				_ = json.Unmarshal([]byte(actual), &actualErr)
			}

			assert.Equal(t, tt.statusCode, response.StatusCode)
			assert.Equal(t, tt.account, actualAccount)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}

func extractResponseFromResponseWriter(responseWriter *httptest.ResponseRecorder) (*http.Response, string) {
	resp := responseWriter.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := string(body)
	return resp, actual
}
