package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "first-topic", 0)
	if err!= nil {
        panic(err)
    }
	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))

	batch := conn.ReadBatch(1e3, 1e9)
	
	bytes := make([]byte, 1e3)

	for  {
		_, err := batch.Read(bytes)
		if err != nil {
            break
        }
		fmt.Println(string(bytes))
	}
}