package sub

import (
	"context"
	"fmt"
	"log"
	"micro-user/utils"

	"github.com/IBM/sarama"
)

type Sub struct {
	Consumer sarama.ConsumerGroup
	Topics   []string
}

func NewSub(topics []string) *Sub {
	return &Sub{
		Consumer: utils.KafkaConsumer,
		Topics:   topics,
	}
}

func (m *Sub) Msg() {
	for {
		err := m.Consumer.Consume(context.Background(), m.Topics, m)
		if err != nil {
			log.Fatalf("error in consuming msg: %v", err)
		}
	}
}

func (m *Sub) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Sub) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Sub) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Println(msg.Topic)
		fmt.Println(msg.Value)
		// var userMessage UserMessage
		// err := json.Unmarshal(message.Value, &userMessage)
		// if err != nil {
		// 	log.Printf("Error unmarshalling message: %v", err)
		// 	continue
		// }
		// log.Printf("Received message: %+v", userMessage)
		// session.MarkMessage(message, "")
	}
	return nil
}
