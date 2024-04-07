package service

import "money_transfer/service/transaction"

type Services struct {
	TransactionService transaction.TransactionService
}

type NewServicesOptions struct {
	TransactionService transaction.TransactionService
}

func NewServices(opts NewServicesOptions) (*Services, error) {
	return &Services{
		TransactionService: opts.TransactionService,
	}, nil
}
