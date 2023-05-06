package event

import "time"

type BalanceUpdated struct {
	Name    string
	Payload interface{}
}

func NewBalanceUpdated() *BalanceUpdated {
	return &BalanceUpdated{
		Name: "BalanceUpdated",
	}
}

func (bu *BalanceUpdated) GetName() string {
	return bu.Name
}

func (bu *BalanceUpdated) GetPayload() interface{} {
	return bu.Payload
}

func (bu *BalanceUpdated) SetPayload(payload interface{}) {
	bu.Payload = payload
}

func (bu *BalanceUpdated) GetDateTime() time.Time {
	return time.Now()
}
