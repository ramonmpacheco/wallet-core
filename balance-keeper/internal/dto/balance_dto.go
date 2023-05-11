package dto

import "time"

type BalanceDTO struct {
	Id        string    `json:"id"`
	AccountId string    `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
