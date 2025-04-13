package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("❌ Error reading config file:", err)
		return
	}

	brokerAddress := viper.GetString("KAFKA_BROKER")
	topic := viper.GetString("KAFKA_TOPIC")

	fmt.Println("📡 Connecting to Kafka broker at: ", brokerAddress)
	fmt.Println(" and subscribing to topic: ", topic)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	fmt.Println("📥 Listening for messages on topic:", topic)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("❌ Read error:", err)
			break
		}
		fmt.Printf("📨 Received [%s] => %s\n", string(msg.Key), string(msg.Value))
	}
}
