package controllers

import (
	"myproject/config"
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTransactionsController(c echo.Context) error {
	transactions := model.Transactions{}
	c.Bind(&transactions)

	err := config.DB.Save(&transactions).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage":      "success create transactions",
		"transactions": transactions,
	})
}

func GetTransactionsController(c echo.Context) error {
	var transactions []model.Transactions

	err := config.DB.Find(&transactions).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactions,
	})
}

func DeleteTransactionsController(c echo.Context) error {
	stringId := c.Param("id")
	err := config.DB.Delete(&model.Transactions{}, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaction with id `" + stringId + "`",
	})
}

func UpdateTransactionsController(c echo.Context) error {
	transaction := model.Transactions{}
	c.Bind(&transaction)
	stringId := c.Param("id")
	err := config.DB.Model(&transaction).Where("id = ?", stringId).Updates(transaction).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success",
		"transaction": transaction,
	})
}
