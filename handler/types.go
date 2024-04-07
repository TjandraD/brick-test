package handler

type ValidateAccountRequest struct {
	AccountNumber int64  `json:"account_number"`
	Name          string `json:"name"`
}

type ValidateAccountResponse struct {
	IsExists bool `json:"is_exists"`
}
