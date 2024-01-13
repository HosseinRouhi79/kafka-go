package main

import (
	"context"
	"fmt"
	"kafgo/logger"
	"log"
	"github.com/segmentio/kafka-go"
)

// read from kafka topic first-topic and all partitions
func main() {
	var logger = logger.Logger()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"}, //we can read from many brokers here
		Topic:     "event-stream",
		GroupID: "all-partitions-group",
		MaxBytes:  10e6,
	})
	logger.Info().Msg("----------starting to read from kafka topic----------")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message is: %s , partition is: %d\n", string(m.Value), m.Partition)
		str := string(m.Value)
		logger.Info().Msg(str)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
	logger.Info().Msg("----------finished reading from kafka topic----------")

}
