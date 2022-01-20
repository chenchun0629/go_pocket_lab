package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"net"
)

func Register(ctx context.Context, client *clientv3.Client, service string, listener net.Listener) error {
	grant, err := client.Grant(ctx, 2)
	if err != nil {
		return err
	}

	kaCh, err := client.KeepAlive(ctx, grant.ID)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-kaCh:
			}
		}
	}()

	em, err := endpoints.NewManager(client, service)
	if err != nil {
		return err
	}

	var ip, _ = Extract(listener.Addr().String(), listener)
	var key = fmt.Sprintf("%s/%s", service, ip)
	err = em.AddEndpoint(
		context.TODO(),
		key,
		endpoints.Endpoint{Addr: ip},
		clientv3.WithLease(grant.ID),
	)
	if err != nil {
		return err
	}

	return nil
}
