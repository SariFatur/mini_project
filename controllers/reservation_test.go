package controllers

import (
	"myproject/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestReservation() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetReservationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get reservation normal",
			path:       "/reservation",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestReservation()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetReservationController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestCreateReservationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "create reservation normal",
			path:       "/reservation",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestReservation()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateReservationController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteReservationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "delete reservation normal",
			path:       "/reservation",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestReservation()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeleteReservationController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateReservationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "update reservation normal",
			path:       "/reservation",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestReservation()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdateReservationController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
