package gateway

import "github.com/ramonmpacheco/ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
