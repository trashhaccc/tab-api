package handler

import (
	"net/http"

	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var income model.Transaction
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.transactionService.CreateTransaction(&income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil menambahkan transaksi"})
}

func (h *TransactionHandler) HistoryTransaction(c *gin.Context) {
	nik := c.Param("nik")
	if nik == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NIK must be not empty"})
	}
	history, err := h.transactionService.HistoryTransaction(nik)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}
