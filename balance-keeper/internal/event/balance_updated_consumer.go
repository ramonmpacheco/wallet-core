package event

import (
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ramonmpacheco/balance-keeper/internal/dto"
	"github.com/ramonmpacheco/balance-keeper/internal/usecase"
	"github.com/ramonmpacheco/balance-keeper/pkg/kafka"
)

type BalanceUpdatedConsumer struct {
	UpdateBalanceUseCase usecase.UpdateBalanceUseCase
}

func NewBalanceUpdatedConsumer(UpdateBalanceUseCase usecase.UpdateBalanceUseCase) *BalanceUpdatedConsumer {
	return &BalanceUpdatedConsumer{
		UpdateBalanceUseCase: UpdateBalanceUseCase,
	}
}

func (e *BalanceUpdatedConsumer) Exec() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"client.id":         "balance-consumer",
		"group.id":          "balance-group",
	}
	topic := []string{"balances"}

	kafkaConsumer := kafka.NewConsumer(&configMap, topic)
	cChannel := make(chan *ckafka.Message)

	log.Default().Println("Start kafka consumer")
	go kafkaConsumer.Consume(cChannel)

	log.Default().Println("Waiting messages")
	for event := range cChannel {
		log.Default().Println("Message found see below:")
		eventValue := dto.EventValue{}
		err := json.Unmarshal(event.Value, &eventValue)
		if err != nil {
			log.Printf("Failed to unmarshal Kafka message: %v", err)
		}
		e.UpdateBalanceUseCase.Execute(eventValue.Payload)
	}
}
