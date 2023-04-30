package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramonmpacheco/ms-wallet/internal/database"
	"github.com/ramonmpacheco/ms-wallet/internal/event"
	createaccount "github.com/ramonmpacheco/ms-wallet/internal/usecase/create_account"
	createclient "github.com/ramonmpacheco/ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/ramonmpacheco/ms-wallet/internal/usecase/create_transaction"
	"github.com/ramonmpacheco/ms-wallet/internal/web"
	"github.com/ramonmpacheco/ms-wallet/internal/web/server"
	"github.com/ramonmpacheco/ms-wallet/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"mysql",
		"3306",
		"wallet",
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTRansactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)
	clientDb := database.NewClientDb(db)
	accounttDb := database.NewAccountDb(db)
	transationDb := database.NewTransactionDb(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accounttDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(
		transationDb, accounttDb, eventDispatcher, transactionCreatedEvent,
	)
	webserver := server.NewWebServer(":3000")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
