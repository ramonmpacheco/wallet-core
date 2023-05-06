package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *ClientDb
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients(id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDb = NewClientDb(db)
}

func (s *ClientDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) Test_Get() {
	client, _ := entity.NewClient("John Doe", "j@d.com")
	s.clientDb.Save(client)

	clientDb, err := s.clientDb.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Name, clientDb.Name)
	s.Equal(client.Email, clientDb.Email)
}

func (s *ClientDbTestSuite) Test_Save() {
	client := &entity.Client{
		ID:    "1",
		Name:  "John Doe",
		Email: "j@d.com",
	}
	err := s.clientDb.Save(client)
	s.Nil(err)
}
