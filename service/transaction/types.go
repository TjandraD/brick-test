package transaction

type CreateTransactionInput struct {
	AccountNumber int64
	Amount        int64
}

type CreateTransactionApiRequest struct {
	AccountNumber int64 `json:"account_number"`
	Amount        int64 `json:"amount"`
}
