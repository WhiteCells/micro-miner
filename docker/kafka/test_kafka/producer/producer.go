package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{
		"localhost:19091",
		"localhost:19092",
		"localhost:19093",
	}
	topic := "test-topic"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("NewSyncProducer error: %v", err)
	}
	defer producer.Close()

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("send message num: %d", i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("producer.SendMessage error: %v", err)
		} else {
			log.Printf("producer.SendMessage success partition=%d, offset=%d", partition, offset)
		}
	}
}
