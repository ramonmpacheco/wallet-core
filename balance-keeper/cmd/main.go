package main

import (
	"log"

	"github.com/ramonmpacheco/balance-keeper/internal/event"
	"github.com/ramonmpacheco/balance-keeper/internal/web"
	"github.com/ramonmpacheco/balance-keeper/internal/web/server"
)

func main() {
	log.Default().Println("Init app...")
	go event.NewBalanceUpdatedConsumer().Exec()
	webserver := server.NewWebServer(":3003")
	balanceHandler := web.NewBalanceHandler()
	webserver.AddHandler("/balances/{account_id}", balanceHandler.GetBalance)
	webserver.Start()
}
