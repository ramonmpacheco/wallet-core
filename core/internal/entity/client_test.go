package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_New_Client(t *testing.T) {
	client, err := NewClient("John Doe", "j@d.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@d.com", client.Email)
}

func Test_Create_New_Test_When_Args_Are_Invalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func Test_Update_Client(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	err := client.Update("John Doe Jr", "j@dj.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Jr", client.Name)
	assert.Equal(t, "j@dj.com", client.Email)
}

func Test_Update_Client_With_Invalid_Args(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	err := client.Update("", "j@dj.com")
	assert.Error(t, err, "name is required")
}

func Test_Add_Account(t *testing.T) {
	client, _ := NewClient("John Doe", "j@d.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
