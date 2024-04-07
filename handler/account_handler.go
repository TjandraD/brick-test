package handler

import (
	"errors"
	"fmt"
	errhelper "money_transfer/common/err_helper"
	"money_transfer/service"
	"money_transfer/service/account"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateAccount(services *service.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ValidateAccountRequest
		if err := c.Bind(&req); err != nil {
			fmt.Printf("error binding request: %s", err.Error())
			return errhelper.ReturnError(c, errors.Join(err))
		}

		validateInput := account.ValidateAccountInput{
			AccountNumber: req.AccountNumber,
			Name:          req.Name,
		}
		isExists, err := services.AccountService.ValidateAccount(validateInput)
		if err != nil {
			fmt.Printf("error validating account: %s", err.Error())
			return errhelper.ReturnError(c, err)
		}

		return c.JSON(http.StatusOK, ValidateAccountResponse{
			IsExists: isExists,
		})
	}
}
