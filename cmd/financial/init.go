package main

import (
	"fmt"

	"github.com/murilosrg/financial-api/internal/database"
	"github.com/murilosrg/financial-api/internal/model"
)

func initAll() {
	drop()
	seed()
}

func drop() {
	if (database.DB.HasTable(&model.Transaction{})) {
		fmt.Println("dropping table transaction.")
		database.DB.DropTable(&model.Transaction{})
	}

	if (database.DB.HasTable(&model.Account{})) {
		fmt.Println("dropping table account.")
		database.DB.DropTable(&model.Account{})
	}

	if (database.DB.HasTable(&model.OperationType{})) {
		fmt.Println("dropping table operation_type.")
		database.DB.DropTable(&model.OperationType{})
	}
}

func seed() {
	database.DB.AutoMigrate(&model.Account{})
	database.DB.AutoMigrate(&model.OperationType{})

	database.DB.AutoMigrate(&model.Transaction{}).
		AddForeignKey("account_id", "account(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("operation_type_id", "operation_type(id)", "RESTRICT", "RESTRICT")

	ot0 := model.OperationType{ID: 1, Description: "COMPRA A VISTA", IsDebit: true}
	ot1 := model.OperationType{ID: 2, Description: "COMPRA PARCELADA", IsDebit: true}
	ot2 := model.OperationType{ID: 3, Description: "SAQUE", IsDebit: true}
	ot3 := model.OperationType{ID: 4, Description: "PAGAMENTO", IsDebit: false}
	_, _ = ot0.Create()
	_, _ = ot1.Create()
	_, _ = ot2.Create()
	_, _ = ot3.Create()
}
