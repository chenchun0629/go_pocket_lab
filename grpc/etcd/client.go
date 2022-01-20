package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func NewClient() *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:23791", "localhost:23792", "localhost:23793"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return client
}
