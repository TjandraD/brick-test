package main

import (
	"fmt"
	"money_transfer/app"
	"money_transfer/handler"
	"money_transfer/repo"
	"money_transfer/service"
	"money_transfer/service/account"
	"money_transfer/service/transaction"

	"github.com/labstack/echo/v4"
)

func main() {
	config := app.LoadConfig()
	db, err := app.InitDb(config)
	if err != nil {
		fmt.Printf("error initializing db: %s", err.Error())
		return
	}

	repo := repo.NewRepository(db)
	servicesOpts := service.NewServicesOptions{
		TransactionService: transaction.NewTransactionService(repo, config),
		AccountService:     account.NewAccountService(repo, config),
	}
	services, err := service.NewServices(servicesOpts)
	if err != nil {
		fmt.Printf("error initializing: %s", err.Error())
		return
	}

	e := echo.New()
	registerHandlers(e, services)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.AppPort)))
}

func registerHandlers(e *echo.Echo, services *service.Services) {
	e.POST("/accounts/validate", handler.ValidateAccount(services))
	e.POST("/transactions", handler.CreateTransaction(services))
	e.POST("/transactions/:id/confirm", handler.ConfirmTransaction(services))
}
