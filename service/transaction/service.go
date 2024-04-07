package transaction

import (
	"bytes"
	"encoding/json"
	"errors"
	"money_transfer/app"
	"money_transfer/constant"
	"money_transfer/repo"
	"net/http"
)

type TransactionServiceInterface interface {
	CreateTransaction(input CreateTransactionInput) (bool, error)
}

type TransactionService struct {
	Repository repo.RepositoryInterface
	Config     app.Config
}

func NewTransactionService(repo repo.RepositoryInterface, config app.Config) TransactionService {
	return TransactionService{
		Repository: repo,
		Config:     config,
	}
}

func (s TransactionService) CreateTransaction(input CreateTransactionInput) (bool, error) {
	isSuccess := false
	transactionInput := repo.Transaction{
		RecipientAccountNumber: input.AccountNumber,
		Amount:                 input.Amount,
		Status:                 constant.TransactionPendingStatus,
	}
	transactionOutput, err := s.Repository.CreateTransaction(transactionInput)
	if err != nil {
		return isSuccess, err
	}

	reqBody := CreateTransactionApiRequest{
		AccountNumber: input.AccountNumber,
		Amount:        input.Amount,
	}

	var bufBody bytes.Buffer
	err = json.NewEncoder(&bufBody).Encode(reqBody)
	if err != nil {
		return isSuccess, errors.Join(constant.ErrBindRequest)
	}

	resp, err := http.Post(s.Config.FinanceApiDomain+"/transaction", "application/json", &bufBody)
	if err != nil {
		return isSuccess, err
	}

	_, err = s.Repository.UpdateTransactionStatus(transactionOutput.ID, constant.TransactionProcessedStatus)
	if err != nil {
		return isSuccess, err
	}

	if resp.StatusCode == http.StatusCreated {
		isSuccess = true
	}

	return isSuccess, nil
}
