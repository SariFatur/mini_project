package controllers

import (
	"myproject/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestTransactions() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetTransactionsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get transaction normal",
			path:       "/transaction",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestTransactions()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetTransactionsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestCreateTransactionsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "create transaction normal",
			path:       "/transaction",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestTransactions()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateTransactionsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteTransactionsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "delete transaction normal",
			path:       "/transaction",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestTransactions()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeleteTransactionsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateTransactionsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "update transaction normal",
			path:       "/transaction",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestTransactions()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdateTransactionsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
