package account

type ValidateAccountInput struct {
	AccountNumber int64
	Name          string
}

type ValidateApiInput struct {
	Name          string `json:"name"`
	AccountNumber int64  `json:"account_number"`
}
