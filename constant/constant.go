package constant

import "errors"

const (
	ErrMsgUnknownError = "Unknown Error"
)

var (
	ErrInvalidRequest       = errors.New("invalid request")
	ErrBindRequest          = errors.New("error binding request")
	ErrInvalidTransactionId = errors.New("invalid transaction id")
)

var (
	TransactionPendingStatus   = "PENDING"
	TransactionProcessedStatus = "PROCESSED"
	TransactionSuccessStatus   = "SUCCESS"
	TransactionFailedStatus    = "FAILED"
)
