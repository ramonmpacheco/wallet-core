package usecase

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ramonmpacheco/balance-keeper/internal/db"
	"github.com/ramonmpacheco/balance-keeper/internal/dto"
	"github.com/ramonmpacheco/balance-keeper/internal/entity"
)

type UpdateBalanceUseCase struct {
	Persistence db.BalanceDb
}

func NewUpdateBalanceUseCase(persistence db.BalanceDb) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		Persistence: persistence,
	}
}

func (du *UpdateBalanceUseCase) Execute(bu dto.BalanceUpdatedDTO) {
	log.Default().Printf("Init usecase: %v", bu)
	du.save(
		du.getEntitiesFrom(bu),
	)
	log.Default().Printf("usecase done: %v", bu)
}

func (s *UpdateBalanceUseCase) save(entities []entity.Balance) {
	for _, entity := range entities {
		e, _ := s.Persistence.GetByAccountId(entity.AccountId)
		if e != nil {
			e.Amount = entity.Amount
			e.UpdatedAt = time.Now()
			s.Persistence.Update(e)
			continue
		}
		if err := s.Persistence.Save(&entity); err != nil {
			log.Default().Printf("error from persistence: %v", err)
			continue
		}
		log.Default().Printf("persistence ok: %v", entity)
	}
}

func (gef *UpdateBalanceUseCase) getEntitiesFrom(bu dto.BalanceUpdatedDTO) []entity.Balance {
	return []entity.Balance{
		{
			Id:        uuid.New().String(),
			AccountId: bu.AccountIdFrom,
			Amount:    bu.BalanceAccountIdFrom,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        uuid.New().String(),
			AccountId: bu.AccountIdTo,
			Amount:    bu.BalanceAccountIdTo,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
