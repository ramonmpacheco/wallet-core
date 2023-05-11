package usecase

import (
	"log"

	"github.com/ramonmpacheco/balance-keeper/internal/db"
	"github.com/ramonmpacheco/balance-keeper/internal/dto"
)

type FindBalanceUseCase struct {
	Persistence db.BalanceDb
}

func NewFindBalanceUseCase(persistence db.BalanceDb) *FindBalanceUseCase {
	return &FindBalanceUseCase{
		Persistence: persistence,
	}
}

func (gbai *FindBalanceUseCase) GetByAccountId(acountId string) (*dto.BalanceDTO, error) {
	log.Default().Printf("Init FindBalanceUseCase: %s", acountId)
	entity, err := gbai.Persistence.GetByAccountId(acountId)
	if err != nil {
		log.Default().Printf("FindBalanceUseCase, error from get : %s", acountId)
		return nil, err
	}
	log.Default().Printf("FindBalanceUseCase done: %v", entity)
	return &dto.BalanceDTO{
		Id:        entity.Id,
		AccountId: entity.AccountId,
		Amount:    entity.Amount,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}
