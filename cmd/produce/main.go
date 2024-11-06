package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	address := "localhost:9092"
	topic := "my.topic"

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": address})

	if err != nil {
		log.Panic("failed to connect to kafka", err)
	}

	log.Println("connected to kafka")

	defer producer.Close()

	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(word),
		}, nil)
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(`{"time":1497014222380,"id":18,"itemid":"Item_184","address":{"city":"Mountain View","state":"CA","zipcode":94041}}`),
	}, nil)
}
