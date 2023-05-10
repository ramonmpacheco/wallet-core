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
	fromEntity := entity.Balance{
		Id:        uuid.New().String(),
		AccountId: bu.AccountIdFrom,
		Amount:    bu.BalanceAccountIdFrom,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	savedFrom, _ := du.Persistence.GetByAccountId(fromEntity.AccountId)
	if savedFrom != nil {
		savedFrom.Amount = fromEntity.Amount
		savedFrom.UpdatedAt = time.Now()
		du.Persistence.Update(savedFrom)
	} else {
		if err := du.Persistence.Save(&fromEntity); err != nil {
			log.Default().Printf("error from persistence: %v", err)
		}
	}
	toEntity := entity.Balance{
		Id:        uuid.New().String(),
		AccountId: bu.AccountIdTo,
		Amount:    bu.BalanceAccountIdTo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	savedTo, _ := du.Persistence.GetByAccountId(toEntity.AccountId)
	if savedTo != nil {
		savedTo.Amount = toEntity.Amount
		savedTo.UpdatedAt = time.Now()
		du.Persistence.Update(savedTo)
	} else {
		if err := du.Persistence.Save(&toEntity); err != nil {
			log.Default().Printf("error from persistence: %v", err)
		}
	}

}
