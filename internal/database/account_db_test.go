package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDb *AccountDb
	client    *entity.Client
}

func (s *AccountDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients(id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts(id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.accountDb = NewAccountDb(db)
	s.client, _ = entity.NewClient("John Doe", "j@d.com")
}
func (s *AccountDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func Test_AccountDb_Test_Suite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) Test_Save() {
	account := entity.NewAccount(s.client)
	err := s.accountDb.Save(account)
	s.Nil(err)
}
func (s *AccountDbTestSuite) Test_Find_By_Id() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)
	account := entity.NewAccount(s.client)
	err := s.accountDb.Save(account)
	s.Nil(err)
	saveAcc, err := s.accountDb.FindById(account.ID)
	s.Nil(err)
	s.Equal(account.ID, saveAcc.ID)
	s.Equal(account.Client.ID, saveAcc.Client.ID)
	s.Equal(account.Balance, saveAcc.Balance)
}
