package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BalanceHandler struct {
}

func NewBalanceHandler() *BalanceHandler {
	return &BalanceHandler{}
}

func (gb *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "account_id")
	log.Default().Println("o id é ", accountId)
}
