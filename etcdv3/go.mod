module pocket_lab/etcdv3

go 1.16

replace (
	google.golang.org/grpc => github.com/grpc/grpc-go v1.43.0
)

require (
	github.com/pkg/errors v0.9.1
	go.etcd.io/etcd/api/v3 v3.5.1
	go.etcd.io/etcd/client/v3 v3.5.1
	google.golang.org/grpc v1.38.0
)
