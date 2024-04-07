package main

import (
	"fmt"
	"money_transfer/app"
	"money_transfer/repo"
	"money_transfer/service"
	"money_transfer/service/transaction"
	"net/http"

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
		TransactionService: transaction.NewTransactionService(repo),
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
