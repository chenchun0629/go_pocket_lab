package main

import (
	"context"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"log"
	"net"
	"pocket_lab/grpc/etcd"
	"pocket_lab/grpc/message"
	"strings"
)

type MessageSender struct {
	message.UnsafeMessageSenderServer
}

func (m *MessageSender) Send(ctx context.Context, request *message.MessageRequest) (*message.MessageResponse, error) {
	log.Println("receive message:", request.GetSaySomething())
	resp := &message.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}

func Register(ctx context.Context, client *clientv3.Client, service, self string) error {
	resp, err := client.Grant(ctx, 2)
	if err != nil {
		return errors.Wrap(err, "etcd grant")
	}
	_, err = client.Put(ctx, strings.Join([]string{service, self}, "/"), self, clientv3.WithLease(resp.ID))
	if err != nil {
		return errors.Wrap(err, "etcd put")
	}
	// respCh 需要消耗, 不然会有 warning
	respCh, err := client.KeepAlive(ctx, resp.ID)
	if err != nil {
		return errors.Wrap(err, "etcd keep alive")
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-respCh:

		}
	}
}

//func etcdRegister(listener net.Listener) {
//
//
//	em, err := endpoints.NewManager(client, "test/grpc/message")
//	if err != nil {
//		panic(err)
//	}
//
//	grant, err := client.Grant(context.TODO(), 2)
//	if err != nil {
//		panic(err)
//	}
//
//	go func() {
//		respCh, err := client.KeepAlive(context.TODO(), grant.ID)
//		if err != nil {
//			panic(err)
//		}
//
//		for {
//			select {
//			case <-respCh:
//			}
//		}
//	}()
//
//	var ip, _ = etcd.Extract(listener.Addr().String(), listener)
//
//	err = em.AddEndpoint(
//		context.TODO(),
//		"test/grpc/message/e1",
//		//endpoints.Endpoint{Addr: "127.0.0.1:12345"},
//		endpoints.Endpoint{Addr: ip},
//		clientv3.WithLease(grant.ID),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	//go func() {
//	//	_ = Register(context.TODO(), client, "", "")
//	//}()
//}

func main() {
	srv := grpc.NewServer()
	message.RegisterMessageSenderServer(srv, &MessageSender{})
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = etcd.Register(context.TODO(), etcd.NewClient(), "test/grpc/message", listener)
	if err != nil {
		log.Fatalf("failed to register serve: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleSendMessage(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	log.Println("receive message:", req.GetSaySomething())
	resp := &message.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}
