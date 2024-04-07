package handler

type ValidateAccountRequest struct {
	AccountNumber int64  `json:"account_number"`
	Name          string `json:"name"`
}

type ValidateAccountResponse struct {
	IsExists bool `json:"is_exists"`
}

type CreateTransactionRequest struct {
	AccountNumber int64 `json:"account_number"`
	Amount        int64 `json:"amount"`
}

type CreateTransactionResponse struct {
	IsSuccess bool `json:"is_success"`
}

type ConfirmTransactionRequest struct {
	IsSuccess bool `json:"is_success"`
}
