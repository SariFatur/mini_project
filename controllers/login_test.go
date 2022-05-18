package controllers

import (
	"myproject/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestLogin() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestLoginController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "login normal",
			path:       "/login",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}
	e := InitEchoTestLogin()
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
