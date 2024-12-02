package service

import (
	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
)

type TransactionService interface {
	CreateTransaction(income *model.Transaction) error
	HistoryTransaction(nik string) ([]model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo}
}

func (s *transactionService) CreateTransaction(income *model.Transaction) error {
	return s.transactionRepo.CreateTransaction(income)
}
func (s *transactionService) HistoryTransaction(nik string) ([]model.Transaction, error) {
	return s.transactionRepo.HistoryTransaction(nik)
}
