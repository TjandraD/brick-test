package handler

import (
	"errors"
	"fmt"
	errhelper "money_transfer/common/err_helper"
	"money_transfer/service"
	"money_transfer/service/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTransaction(services *service.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateTransactionRequest
		if err := c.Bind(&req); err != nil {
			fmt.Printf("error binding request: %s", err.Error())
			return errhelper.ReturnError(c, errors.Join(err))
		}

		transactionInput := transaction.CreateTransactionInput{
			AccountNumber: req.AccountNumber,
			Amount:        req.Amount,
		}
		isSuccess, err := services.TransactionService.CreateTransaction(transactionInput)
		if err != nil {
			fmt.Printf("error creating transaction: %s", err.Error())
			return errhelper.ReturnError(c, err)
		}

		return c.JSON(http.StatusCreated, CreateTransactionResponse{
			IsSuccess: isSuccess,
		})
	}
}

func ConfirmTransaction(services *service.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionId := c.Param("id")
		var req ConfirmTransactionRequest
		if err := c.Bind(&req); err != nil {
			fmt.Printf("error binding request: %s", err.Error())
			return errhelper.ReturnError(c, errors.Join(err))
		}

		err := services.TransactionService.ConfirmTransaction(transactionId, req.IsSuccess)
		if err != nil {
			fmt.Printf("error confirming transaction: %s", err.Error())
			return errhelper.ReturnError(c, err)
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
