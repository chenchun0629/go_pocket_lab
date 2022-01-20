package main

import (
	"context"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"strings"
	"time"
)

const (
	// etcd resolver 负责的 scheme 类型
	Scheme      = "etcd"
	defaultFreq = time.Minute * 30
)

type Builder struct {
	client *clientv3.Client
	// 全局路由表快照, 非必要
	store map[string]map[string]struct{}
}

func NewBuilder(client *clientv3.Client) *Builder {
	return &Builder{
		client: client,
		store:  make(map[string]map[string]struct{}),
	}
}

func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	b.store[target.Endpoint] = make(map[string]struct{})

	// 初始化 etcd resolver
	r := &etcdResolver{
		client: b.client,
		target: target,
		cc:     cc,
		store:  b.store[target.Endpoint],
		stopCh: make(chan struct{}, 1),
		rn:     make(chan struct{}, 1),
		t:      time.NewTicker(defaultFreq),
	}

	// 开启后台更新 goroutine
	go r.start(context.Background())
	// 全量更新服务地址
	r.ResolveNow(resolver.ResolveNowOptions{})

	return r, nil
}

func (b *Builder) Scheme() string {
	return Scheme
}

type etcdResolver struct {
	client *clientv3.Client
	target resolver.Target
	cc     resolver.ClientConn
	store  map[string]struct{}
	stopCh chan struct{}
	// rn channel is used by ResolveNow() to force an immediate resolution of the target.
	rn chan struct{}
	t  *time.Ticker
}

func (r *etcdResolver) start(ctx context.Context) {
	target := r.target.Endpoint

	w := clientv3.NewWatcher(r.client)
	rch := w.Watch(ctx, target+"/", clientv3.WithPrefix())
	for {
		select {
		case <-r.rn:
			r.resolveNow()
		case <-r.t.C:
			r.ResolveNow(resolver.ResolveNowOptions{})
		case <-r.stopCh:
			w.Close()
			return
		case wresp := <-rch:
			for _, ev := range wresp.Events {
				switch ev.Type {
				case mvccpb.PUT:
					r.store[string(ev.Kv.Value)] = struct{}{}
				case mvccpb.DELETE:
					delete(r.store, strings.Replace(string(ev.Kv.Key), target+"/", "", 1))
				}
			}
			r.updateTargetState()
		}
	}
}

func (r *etcdResolver) resolveNow() {
	target := r.target.Endpoint
	resp, err := r.client.Get(context.Background(), target+"/", clientv3.WithPrefix())
	if err != nil {
		r.cc.ReportError(errors.Wrap(err, "get init endpoints"))
		return
	}

	for _, kv := range resp.Kvs {
		r.store[string(kv.Value)] = struct{}{}
	}

	r.updateTargetState()
}

func (r *etcdResolver) updateTargetState() {
	addrs := make([]resolver.Address, len(r.store))
	i := 0
	for k := range r.store {
		addrs[i] = resolver.Address{Addr: k}
		i++
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

// 会并发调用, 所以这里防止同时多次全量刷新
func (r *etcdResolver) ResolveNow(o resolver.ResolveNowOptions) {
	select {
	case r.rn <- struct{}{}:
	default:

	}
}

func (r *etcdResolver) Close() {
	r.t.Stop()
	close(r.stopCh)
}
