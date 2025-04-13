package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("❌ Error reading config file:", err)
		return
	}

	brokerAddress := viper.GetString("KAFKA_BROKER")
	topic := viper.GetString("KAFKA_TOPIC")

	fmt.Println("📡 Connecting to Kafka broker at: ", brokerAddress)
	fmt.Println(" and publishing to topic: ", topic)

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	message := "Hello Kafka from Golang!"
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", time.Now().Unix())),
			Value: []byte(message),
		},
	)
	if err != nil {
		fmt.Println("❌ Failed to write message:", err)
		return
	}

	fmt.Println("✅ Sent:", message)
}
