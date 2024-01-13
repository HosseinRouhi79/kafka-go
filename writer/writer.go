package main

import (
	"context"
	"kafgo/logger"

	"github.com/segmentio/kafka-go"
)

func main() {
	var logger = logger.Logger()

	
	w := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:                  "second-topic",
		AllowAutoTopicCreation: true,
		Balancer: &kafka.LeastBytes{},
	}
	messages := []kafka.Message{
		{
			Key:   []byte("key1"),
            Value: []byte("value1"),
		},
		{
            Key:   []byte("key2"),
            Value: []byte("value2"),
        },
		{
            Key:   []byte("key3"),
            Value: []byte("value3"),
        },
	}
	for i := 0; i < len(messages); i++ {
		w.WriteMessages(context.Background(), messages[i])
	}

	if err := w.Close(); err!= nil {
        logger.Fatal().Msg("failed to close writer:")
    }
	

}
