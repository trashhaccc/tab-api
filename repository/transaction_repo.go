package repository

import (
	"github.com/HIUNCY/rest-api-go/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(income *model.Transaction) error
	HistoryTransaction(nik string) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(income *model.Transaction) error {
	return r.db.Create(income).Error
}

func (r *transactionRepository) HistoryTransaction(nik string) ([]model.Transaction, error) {
	var transaction []model.Transaction
	if err := r.db.Where("nik = ?", nik).Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
