package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Account(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")

	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func Test_Create_Account_With_Client_Nil(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func Test_Credit_Account(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func Test_Debit_Account(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
