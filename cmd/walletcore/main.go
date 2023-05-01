package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramonmpacheco/ms-wallet/internal/database"
	"github.com/ramonmpacheco/ms-wallet/internal/event"
	"github.com/ramonmpacheco/ms-wallet/internal/event/handler"
	createaccount "github.com/ramonmpacheco/ms-wallet/internal/usecase/create_account"
	createclient "github.com/ramonmpacheco/ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/ramonmpacheco/ms-wallet/internal/usecase/create_transaction"
	"github.com/ramonmpacheco/ms-wallet/internal/web"
	"github.com/ramonmpacheco/ms-wallet/internal/web/server"
	"github.com/ramonmpacheco/ms-wallet/pkg/events"
	"github.com/ramonmpacheco/ms-wallet/pkg/kafka"
	"github.com/ramonmpacheco/ms-wallet/pkg/uow"
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
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTRansactionCreated()

	clientDb := database.NewClientDb(db)
	accounttDb := database.NewAccountDb(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)
	uow.Register("AccountDb", func(tx *sql.Tx) interface{} {
		return database.NewAccountDb(db)
	})
	uow.Register("TransactionDb", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDb(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accounttDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(
		uow, eventDispatcher, transactionCreatedEvent,
	)
	webserver := server.NewWebServer(":8080")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
