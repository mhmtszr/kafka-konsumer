package main

import (
	"fmt"
	"github.com/Trendyol/kafka-konsumer"
)

func main() {
	consumerCfg := &kafka.ConsumerConfig{
		Concurrency: 1,
		Reader: kafka.ReaderConfig{
			Brokers: []string{"localhost:29092"},
			Topic:   "standart-topic",
			GroupID: "standart-cg",
		},
		RetryEnabled: false,
		ConsumeFn:    consumeFn,
	}

	consumer, _ := kafka.NewConsumer(consumerCfg)
	defer consumer.Stop()

	consumer.Consume()

	fmt.Println("Consumer started...!")
	select {}
}

func consumeFn(message kafka.Message) error {
	fmt.Printf("Message From %s with value %s", message.Topic, string(message.Value))
	return nil
}
