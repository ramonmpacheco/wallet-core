package createclient

import (
	"log"
	"time"

	"github.com/ramonmpacheco/ms-wallet/internal/entity"
	"github.com/ramonmpacheco/ms-wallet/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (ccud *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	log.Default().Printf("create-client-use-case-execute, init, input=%s", input)
	err = ccud.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}
	log.Default().Printf("create-client-use-case-execute, done, input=%s", input)
	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
