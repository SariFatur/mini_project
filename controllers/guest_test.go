package controllers

import (
	"myproject/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetGuestController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get guest normal",
			path:       "/guest",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetGuestController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestCreateGuestController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "create guest normal",
			path:       "/guest",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateGuestController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestGetGuestByIdController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get guest by id normal",
			path:       "/guest",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetGuestByIdController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteGuestController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "delete guest normal",
			path:       "/guest",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeleteGuestController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateGuestController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "update guest normal",
			path:       "/guest",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdateGuestController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
