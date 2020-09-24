package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/internal/model/transactions"
)

type ITransactionController interface {
	Post(c *gin.Context)
}

type TransactionController struct {
	service transactions.ITransactionService
}

func NewTransactionController(service transactions.ITransactionService) ITransactionController {
	return &TransactionController{
		service: service,
	}
}

func (t *TransactionController) Post(c *gin.Context) {
	transaction := &transactions.Transaction{}
	c.BindJSON(&transaction)

	transaction, err := t.service.Create(transaction)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
