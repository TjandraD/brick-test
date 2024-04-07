package transaction

import (
	"bytes"
	"encoding/json"
	"errors"
	"money_transfer/app"
	"money_transfer/constant"
	"money_transfer/repo"
	"net/http"
	"strconv"
	"sync"
)

type TransactionServiceInterface interface {
	CreateTransaction(input CreateTransactionInput) (bool, error)
	ConfirmTransaction(transactionId string, isSuccess bool) error
}

type TransactionService struct {
	Repository             repo.RepositoryInterface
	Config                 app.Config
	TransactionLocks       map[int]*sync.Mutex
	TransactionStatusLocks map[string]*sync.Mutex
}

func NewTransactionService(repo repo.RepositoryInterface, config app.Config) TransactionService {
	return TransactionService{
		Repository:             repo,
		Config:                 config,
		TransactionLocks:       map[int]*sync.Mutex{},
		TransactionStatusLocks: map[string]*sync.Mutex{},
	}
}

func (s TransactionService) CreateTransaction(input CreateTransactionInput) (bool, error) {
	if mut, ok := s.TransactionLocks[int(input.AccountNumber)]; ok {
		mut.Lock()
		defer mut.Unlock()
	} else {
		s.TransactionLocks[int(input.AccountNumber)] = &sync.Mutex{}
		s.TransactionLocks[int(input.AccountNumber)].Lock()
		defer s.TransactionLocks[int(input.AccountNumber)].Unlock()
	}

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

func (s TransactionService) ConfirmTransaction(transactionId string, isSuccess bool) error {
	if mut, ok := s.TransactionStatusLocks[transactionId]; ok {
		mut.Lock()
		defer mut.Unlock()
	} else {
		s.TransactionStatusLocks[transactionId] = &sync.Mutex{}
		s.TransactionStatusLocks[transactionId].Lock()
		defer s.TransactionStatusLocks[transactionId].Unlock()
	}

	transactionIdInt, err := strconv.Atoi(transactionId)
	if err != nil {
		return errors.Join(constant.ErrInvalidTransactionId)
	}

	status := constant.TransactionSuccessStatus
	if !isSuccess {
		status = constant.TransactionFailedStatus
	}

	_, err = s.Repository.UpdateTransactionStatus(uint(transactionIdInt), status)
	if err != nil {
		return err
	}

	return nil
}
