package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Transaction(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	client2, _ := NewClient("John Doe 2", "j@d2.com")

	account := NewAccount(client)
	account2 := NewAccount(client2)

	account.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)

	assert.Equal(t, float64(1100), account2.Balance)
	assert.Equal(t, float64(900), account.Balance)
}

func Test_Create_Transaction_With_Insuficient_Funds(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	client2, _ := NewClient("John Doe 2", "j@d2.com")

	account := NewAccount(client)
	account2 := NewAccount(client2)

	account.Credit(50)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 100)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, float64(1000), account2.Balance)
	assert.Equal(t, float64(50), account.Balance)
	assert.Equal(t, "insuficient funds", err.Error())
}
