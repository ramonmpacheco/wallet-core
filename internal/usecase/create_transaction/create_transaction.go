package createtransaction

import (
	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/ramonmpacheco/ms-wallet/internal/gateway"
	"github.com/ramonmpacheco/ms-wallet/pkg/events"
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
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(tg gateway.TransactionGateway, ag gateway.AccountGateway, eventDispatcher events.EventDispatcherInterface, transactionCreated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: tg,
		AccountGateway:     ag,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
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
	output := &CreateTransactionOutputDTO{Id: transaction.ID}
	ctuc.TransactionCreated.SetPayload(output)
	ctuc.EventDispatcher.Dispatch(ctuc.TransactionCreated)
	return output, nil
}
