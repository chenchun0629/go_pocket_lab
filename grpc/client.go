package main

import (
	"context"
	"fmt"
	"log"
	"pocket_lab/grpc/etcd"
	"pocket_lab/grpc/message"
)

func main() {
	etcdcli := etcd.NewClient()
	conn, err := etcd.NewConn(etcdcli, "test/grpc/message")
	if err != nil {
		log.Fatalf("new conn error: %v", err)
	}
	defer conn.Close()

	client := message.NewMessageSenderClient(conn)
	resp, err := client.Send(context.Background(), &message.MessageRequest{SaySomething: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println(resp)
}
