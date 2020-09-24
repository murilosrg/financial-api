package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/internal/controllers"
	"github.com/murilosrg/financial-api/internal/database"
	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/murilosrg/financial-api/internal/model/operations"
	"github.com/murilosrg/financial-api/internal/model/transactions"
)

func SetupApiRouter(r *gin.Engine) {
	account := handleAccount()
	transaction := handleTransaction()

	api := r.Group("/api")

	api.POST("/accounts", account.Post)
	api.GET("/accounts/:accountId", account.Find)
	api.POST("/transactions", transaction.Post)
}

func handleAccount() controllers.IAccountController {
	accountRepo := accounts.NewAccountRepository(database.DB)
	accountService := accounts.NewAccountService(accountRepo)

	return controllers.NewAccountController(accountService)
}

func handleTransaction() controllers.ITransactionController {
	transactionRepo := transactions.NewTransactionRepository(database.DB)
	operationRepo := operations.NewOperationTypeRepository(database.DB)
	accountRepo := accounts.NewAccountRepository(database.DB)
	accountService := accounts.NewAccountService(accountRepo)
	transactionService := transactions.NewTransactionService(accountService, transactionRepo, operationRepo)

	return controllers.NewTransactionController(transactionService)
}
