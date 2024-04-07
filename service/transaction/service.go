package transaction

import "money_transfer/repo"

type TransactionServiceInterface interface{}

type TransactionService struct {
	Repository *repo.Repository
}

func NewTransactionService(repo *repo.Repository) TransactionService {
	return TransactionService{
		Repository: repo,
	}
}
