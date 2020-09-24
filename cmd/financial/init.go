package main

import (
	"fmt"

	"github.com/murilosrg/financial-api/config"
	"github.com/murilosrg/financial-api/internal/database"
	"github.com/murilosrg/financial-api/internal/model/accounts"
	"github.com/murilosrg/financial-api/internal/model/operations"
	"github.com/murilosrg/financial-api/internal/model/transactions"
)

func initAll() {
	drop()
	seed()
}

func drop() {
	if (database.DB.HasTable(&transactions.Transaction{})) {
		fmt.Println("dropping table transaction.")
		database.DB.DropTable(&transactions.Transaction{})
	}

	if (database.DB.HasTable(&accounts.Account{})) {
		fmt.Println("dropping table account.")
		database.DB.DropTable(&accounts.Account{})
	}

	if (database.DB.HasTable(&operations.OperationType{})) {
		fmt.Println("dropping table operation_type.")
		database.DB.DropTable(&operations.OperationType{})
	}
}

func seed() {
	database.DB.AutoMigrate(&accounts.Account{})
	database.DB.AutoMigrate(&operations.OperationType{})

	if config.Load().DB.Driver == "sqlite3" {
		database.DB.AutoMigrate(&transactions.Transaction{})
	} else {
		database.DB.AutoMigrate(&transactions.Transaction{}).
			AddForeignKey("account_id", "account(id)", "RESTRICT", "RESTRICT").
			AddForeignKey("operation_type_id", "operation_type(id)", "RESTRICT", "RESTRICT")
	}

	ot0 := &operations.OperationType{ID: 1, Description: "COMPRA A VISTA", IsDebit: true}
	ot1 := &operations.OperationType{ID: 2, Description: "COMPRA PARCELADA", IsDebit: true}
	ot2 := &operations.OperationType{ID: 3, Description: "SAQUE", IsDebit: true}
	ot3 := &operations.OperationType{ID: 4, Description: "PAGAMENTO", IsDebit: false}

	or := operations.NewOperationTypeRepository(database.DB)
	_ = or.Create(ot0)
	_ = or.Create(ot1)
	_ = or.Create(ot2)
	_ = or.Create(ot3)
}
