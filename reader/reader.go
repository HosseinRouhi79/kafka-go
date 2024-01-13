package main

import (
	"context"
	"fmt"
	"kafgo/logger"
	"log"
	// "os"

	"github.com/segmentio/kafka-go"
)

// read from kafka topic first-topic and partition 0
func main() {
	var logger = logger.Logger()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"}, //we can read from many brokers here
		Topic:     "second-topic",
		Partition: 0,
		MaxBytes:  10e6,
	})
	logger.Info().Msg("----------starting to read from kafka topic----------")

	for i := 0; i < 30 ; i++ {
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
	// os.Exit(401)
	logger.Info().Msg("----------finished reading from kafka topic----------")

}
