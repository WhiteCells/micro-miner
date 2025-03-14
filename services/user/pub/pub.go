package pub

import (
	"encoding/json"
	"log"
	"micro-user/utils"

	"github.com/IBM/sarama"
)

func PubMsg(topic string, value any) {
	valueByte, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("failed to marshal value: %v", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(valueByte),
	}
	_, _, err = utils.KafkaProducer.SendMessage(msg)
	if err != nil {
		log.Fatalf("failed send msg to kafka: %v", err)
	}
	log.Println("msg send success")
}
