package controllers

import (
	"myproject/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestPaymentMethod() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetPaymentMethodController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get PaymentMethod normal",
			path:       "/PM",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestPaymentMethod()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetPaymentMethodController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestCreatePaymentMethodController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "create PaymentMethod normal",
			path:       "/PM",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestPaymentMethod()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreatePaymentMethodController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeletePaymentMethodController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "delete PaymentMethod normal",
			path:       "/PM",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestPaymentMethod()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeletePaymentMethodController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdatePaymentMethodController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "update PaymentMethod normal",
			path:       "/PM",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestPaymentMethod()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdatePaymentMethodController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
