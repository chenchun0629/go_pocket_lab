package etcd

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

func NewConn(etcdcli *clientv3.Client, service string) (*grpc.ClientConn, error) {
	r, _ := resolver.NewBuilder(etcdcli)
	return grpc.Dial(fmt.Sprintf("etcd:///%s", service), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithResolvers(r))
}
