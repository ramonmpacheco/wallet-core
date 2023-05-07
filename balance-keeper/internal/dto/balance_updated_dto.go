package dto

type EventValue struct {
	Name    string            `json:"Name"`
	Payload BalanceUpdatedDTO `json:"Payload"`
}
type BalanceUpdatedDTO struct {
	AccountIdFrom        string  `json:"account_id_from"`
	AccountIdTo          string  `json:"account_id_to"`
	BalanceAccountIdFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIdTo   float64 `json:"balance_account_id_to"`
}
