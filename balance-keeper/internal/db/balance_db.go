package db

import (
	"database/sql"
	"log"

	"github.com/ramonmpacheco/balance-keeper/internal/entity"
)

type BalanceDb struct {
	DB *sql.DB
}

func NewBalanceDb(db *sql.DB) *BalanceDb {
	return &BalanceDb{
		DB: db,
	}
}

func (b *BalanceDb) Save(balance *entity.Balance) error {
	log.Default().Printf("Init save balanceDb: %v", balance)
	stmt, err := b.DB.Prepare("INSERT INTO balances (id, account_id, amount, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Id, balance.AccountId, balance.Amount, balance.CreatedAt, balance.UpdatedAt)
	if err != nil {
		log.Default().Printf("error on saving: %s", err.Error())
		return err
	}
	log.Default().Printf("balance saved: %v", balance)
	return nil
}

func (b *BalanceDb) Update(balance *entity.Balance) error {
	log.Default().Printf("Init update balanceDb: %v", balance)
	stmt, err := b.DB.Prepare("UPDATE balances SET amount = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Amount, balance.UpdatedAt, balance.Id)
	if err != nil {
		log.Default().Printf("error on updating: %s", err.Error())
		return err
	}
	log.Default().Printf("balance updated: %v", balance)
	return nil
}

func (b *BalanceDb) GetByAccountId(id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := b.DB.Prepare("SELECT * FROM balances WHERE account_id = ?")
	if err != nil {
		log.Default().Printf("error on parse sql: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&balance.Id, &balance.AccountId, &balance.Amount, &balance.CreatedAt, &balance.UpdatedAt); err != nil {
		log.Default().Printf("error on row scan: %s", err.Error())
		return nil, err
	}
	return balance, nil
}
