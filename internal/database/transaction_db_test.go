package database

import (
	"database/sql"
	"testing"

	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db            *sql.DB
	accountFrom   *entity.Account
	accountTo     *entity.Account
	client        *entity.Client
	client2       *entity.Client
	transactionDb *TransactionDb
}

func (s *TransactionDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients(id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts(id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	client, err := entity.NewClient("John Doe", "j@d.com")
	s.Nil(err)
	client2, err := entity.NewClient("John Doe 2", "j@d2.com")
	s.Nil(err)
	s.client = client
	s.client2 = client2

	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDb = NewTransactionDb(db)
}
func (s *TransactionDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE transactions")
}

func Test_TransactionDb_Test_Suite(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) Test_Create() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDb.Create(transaction)
	s.Nil(err)
}
