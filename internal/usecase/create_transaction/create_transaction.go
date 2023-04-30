package createtransaction

import (
	"context"
	"log"

	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/ramonmpacheco/ms-wallet/internal/gateway"
	"github.com/ramonmpacheco/ms-wallet/pkg/events"
	"github.com/ramonmpacheco/ms-wallet/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string  `json:"account_id_from"`
	AccountIdTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	Id            string  `json:"id"`
	AccountIdFrom string  `json:"account_id_from"`
	AccountIdTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	Uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                Uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (ctuc *CreateTransactionUseCase) Execute(
	ctx context.Context, input CreateTransactionInputDTO,
) (*CreateTransactionOutputDTO, error) {
	log.Default().Printf("CreateTransactionUseCase-execute, init, input=%v", input)
	output := &CreateTransactionOutputDTO{}
	err := ctuc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := ctuc.getAccountRepository(ctx)
		transactionRepository := ctuc.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindById(input.AccountIdFrom)
		if err != nil {
			log.Default().Printf("CreateTransactionUseCase-execute-find-account-from, error, input=%v", err)
			return err
		}
		accountTo, err := accountRepository.FindById(input.AccountIdTo)
		if err != nil {
			log.Default().Printf("CreateTransactionUseCase-execute-find-account-to, error, input=%v", err)
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			log.Default().Printf("CreateTransactionUseCase-execute-new-transaction, error, input=%v", err)
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			log.Default().Printf("CreateTransactionUseCase-execute-update-balance-from, error, input=%v", err)
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			log.Default().Printf("CreateTransactionUseCase-execute-update-balance-to, error, input=%v", err)
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			log.Default().Printf("CreateTransactionUseCase-execute-create-transaction, error, input=%v", err)
			return err
		}
		output.Id = transaction.ID
		output.AccountIdFrom = input.AccountIdFrom
		output.AccountIdTo = input.AccountIdTo
		output.Amount = input.Amount
		return nil
	})
	if err != nil {
		log.Default().Printf("CreateTransactionUseCase-execute, error, input=%v", err)
		return nil, err
	}
	ctuc.TransactionCreated.SetPayload(output)
	ctuc.EventDispatcher.Dispatch(ctuc.TransactionCreated)
	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDb")
	if err != nil {
		log.Default().Printf("CreateTransactionUseCase-execute-getAccountRepository, error, input=%v", err)
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDb")
	if err != nil {
		log.Default().Printf("CreateTransactionUseCase-execute-getTransactionRepository, error, input=%v", err)
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
