package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ramonmpacheco/balance-keeper/internal/usecase"
)

type BalanceHandler struct {
	FindUsecase *usecase.FindBalanceUseCase
}

func NewBalanceHandler(fuc *usecase.FindBalanceUseCase) *BalanceHandler {
	return &BalanceHandler{
		FindUsecase: fuc,
	}
}

func (gb *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	accountId := chi.URLParam(r, "account_id")
	if accountId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid account_id"))
		return
	}
	log.Default().Printf("Init GetBalance: %s", accountId)
	result, err := gb.FindUsecase.GetByAccountId(accountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Default().Printf("GetBalance done, %v", result)
	w.WriteHeader(http.StatusOK)
}
