package createtransaction

import (
	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/ramonmpacheco/ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	Id string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(tg gateway.TransactionGateway, ag gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: tg,
		AccountGateway:     ag,
	}
}

func (ctuc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := ctuc.AccountGateway.FindById(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := ctuc.AccountGateway.FindById(input.AccountIdTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = ctuc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutputDTO{Id: transaction.ID}, nil
}
