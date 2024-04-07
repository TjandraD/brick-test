package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"money_transfer/app"
	"money_transfer/constant"
	"money_transfer/repo"
	"net/http"
)

type AccountServiceInterface interface {
	ValidateAccount(input ValidateAccountInput) (bool, error)
}

type AccountService struct {
	Repository *repo.Repository
	Config     app.Config
}

func NewAccountService(repo *repo.Repository, config app.Config) AccountService {
	return AccountService{
		Repository: repo,
		Config:     config,
	}
}

func (a AccountService) ValidateAccount(input ValidateAccountInput) (bool, error) {
	isExists := false
	reqBody := ValidateApiInput{
		Name:          input.Name,
		AccountNumber: input.AccountNumber,
	}

	var bufBody bytes.Buffer
	err := json.NewEncoder(&bufBody).Encode(reqBody)
	if err != nil {
		return isExists, errors.Join(constant.ErrBindRequest)
	}

	resp, err := http.Post(a.Config.FinanceApiDomain+"/account", "application/json", &bufBody)
	if err != nil {
		return isExists, err
	}

	if resp.StatusCode == http.StatusCreated {
		isExists = true
	}

	return isExists, nil
}
