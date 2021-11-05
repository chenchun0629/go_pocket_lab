package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	//r := kafka.NewReader(kafka.ReaderConfig{
	//	Brokers:   []string{"localhost:9092"},
	//	Topic:     "my-topic",
	//	Partition: 0,
	//	MinBytes:  10e3, // 10KB
	//	MaxBytes:  10e6, // 10MB
	//})

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   "consumer-group-id",
		Topic:     "my-topic",
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})


	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	//for {
	//	m, err := r.ReadMessage(context.Background())
	//	if err != nil {
	//		break
	//	}
	//	fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	//}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}


func main1() {
	// to consume messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}