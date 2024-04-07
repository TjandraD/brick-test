package service

import (
	"money_transfer/service/account"
	"money_transfer/service/transaction"
)

type Services struct {
	TransactionService transaction.TransactionServiceInterface
	AccountService     account.AccountServiceInterface
}

type NewServicesOptions struct {
	TransactionService transaction.TransactionServiceInterface
	AccountService     account.AccountServiceInterface
}

func NewServices(opts NewServicesOptions) (*Services, error) {
	return &Services{
		TransactionService: opts.TransactionService,
		AccountService:     opts.AccountService,
	}, nil
}
