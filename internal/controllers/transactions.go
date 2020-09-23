package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/internal/model"
)

func PostTransaction(c *gin.Context) {

	transaction := model.Transaction{}
	c.BindJSON(&transaction)

	account := model.Account{}
	if err := account.Find(transaction.AccountID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}

	operation := model.OperationType{}
	if err := operation.Find(transaction.OperationTypeID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "operation invalid"})
		return
	}

	if operation.IsDebit {
		transaction.Amount = transaction.Amount.Neg()
	}

	if err := transaction.Create(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
