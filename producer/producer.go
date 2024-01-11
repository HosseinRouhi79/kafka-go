package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "127.0.0.1:9092", "first-topic", 0)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	conn.WriteMessages(kafka.Message{Value: []byte("the message from kafka producer in golang")})
}
