package service

type TransactionService interface{}

type transactionservice struct{}

type TransactionParams struct {
}

func ProvideService(p TransactionParams) TransactionService {
	return &transactionservice{}
}
