package transaction

import "money_transfer/repo"

type TransactionService struct {
	Repository *repo.Repository
}

func NewTransactionService(repo *repo.Repository) TransactionService {
	return TransactionService{
		Repository: repo,
	}
}
