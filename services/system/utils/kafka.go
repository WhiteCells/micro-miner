package utils

import (
	"log"

	"github.com/IBM/sarama"
)

var KafkaProducer sarama.SyncProducer
var KafkaConsumer sarama.ConsumerGroup

func InitKafka(brokers []string) {
	KafkaProducer = NewKafkaProducer(brokers)
	KafkaConsumer = NewKafkaConsumer(brokers)
}

func NewKafkaProducer(brokers []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("failed to start kafka producer: %v", err)
	}
	log.Println("success to start kafka producer")
	return producer
}

func NewKafkaConsumer(brokers []string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	consumer, err := sarama.NewConsumerGroup(brokers, "microservices-consumer-group", config)
	if err != nil {
		log.Fatalf("failed to start kafka consumer: %v", err)
	}
	log.Println("success to start kafka consumer")
	return consumer
}
