package main

import (
	"context"
	"kafgo/logger"

	"github.com/segmentio/kafka-go"
)

func main() {
	var logger = logger.Logger()

	w := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Balancer: &kafka.CRC32Balancer{},
	}
	messages := []kafka.Message{
		{
			Topic: "first-topic",
			Key:   []byte("test"),
			Value: []byte("value"),
		},
		{
			Topic: "event-stream",
			Key:   []byte("test2"),
			Value: []byte("value"),
		},
	}
	for i := 0; i < len(messages); i++ {
		w.WriteMessages(context.Background(), messages[i])
	}

	if err := w.Close(); err != nil {
		logger.Fatal().Msg("failed to close writer:")
	}

}
