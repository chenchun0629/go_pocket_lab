package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/client/v3"
	"log"
	"net"
	"time"
)

type ServerConfig struct {
	Name string
	Host string
	Port int
}

func (config *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

type Server struct {
	config   *ServerConfig
	listener net.Listener
}

func newServer(config *ServerConfig) (*Server, error) {
	listener, err := net.Listen("tcp", config.Address())
	if err != nil {
		return nil, err
	}
	config.Port = listener.Addr().(*net.TCPAddr).Port
	return &Server{
		config:   config,
		listener: listener,
	}, nil
}

func main() {
	var server, err = newServer(&ServerConfig{
		Name: "test-server",
		Host: "10.10.40.170",
		Port: 8098,
	})

	if err != nil {
		log.Fatal(err)
	}

	RegisterServiceToETCD(server.config.Name, server.listener.Addr().String())
}

func RegisterServiceToETCD(ServiceTarget string, value string) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:23791", "localhost:23792", "localhost:23793"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	var curLeaseId clientv3.LeaseID = 0

	for {
		if curLeaseId == 0 {
			leaseResp, err := lease.Grant(context.TODO(), 10)
			if err != nil {
				panic(err)
			}

			key := ServiceTarget + fmt.Sprintf("%d", leaseResp.ID)
			if _, err := kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
				panic(err)
			}
			curLeaseId = leaseResp.ID
		} else {
			// 续约租约，如果租约已经过期将curLeaseId复位到0重新走创建租约的逻辑
			if _, err := lease.KeepAliveOnce(context.TODO(), curLeaseId); err == rpctypes.ErrLeaseNotFound {
				curLeaseId = 0
				continue
			}
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
