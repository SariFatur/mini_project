package controllers

import (
	"myproject/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestRoom() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetRoomController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get Room normal",
			path:       "/room",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestRoom()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetRoomController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestCreateRoomControllerr(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "create PaymentMethod normal",
			path:       "/room",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestRoom()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateRoomController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteRoomController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "delete Room normal",
			path:       "/room",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestRoom()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeleteRoomController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateRoomController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "update Room normal",
			path:       "/room",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestRoom()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdateRoomController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
