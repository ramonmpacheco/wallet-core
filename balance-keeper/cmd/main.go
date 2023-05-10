package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramonmpacheco/balance-keeper/internal/db"
	"github.com/ramonmpacheco/balance-keeper/internal/event"
	"github.com/ramonmpacheco/balance-keeper/internal/usecase"
	"github.com/ramonmpacheco/balance-keeper/internal/web"
	"github.com/ramonmpacheco/balance-keeper/internal/web/server"
)

func main() {
	log.Default().Println("Init app...")
	database, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"balancekeepermysql",
		"3306",
		"balances",
	))
	if err != nil {
		panic(err)
	}
	defer database.Close()
	balanceDb := db.NewBalanceDb(database)
	usecase := usecase.NewUpdateBalanceUseCase(*balanceDb)
	go event.NewBalanceUpdatedConsumer(*usecase).Exec()
	webserver := server.NewWebServer(":3003")
	balanceHandler := web.NewBalanceHandler()
	webserver.AddHandler("/balances/{account_id}", balanceHandler.GetBalance)
	webserver.Start()
}
