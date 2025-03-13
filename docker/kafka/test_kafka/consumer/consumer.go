package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{
		"localhost:19091",
		"localhost:19092",
		"localhost:19093",
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		println("-")
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d partition %d: %s\n", msg.Offset, msg.Partition, string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v\n", err)
		case <-signals:
			log.Println("Shutting down consumer...")
			return
		}
	}
}
