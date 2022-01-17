package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var conf = clientv3.Config{
		Endpoints:   []string{"localhost:23791", "localhost:23792", "localhost:23793"},
		DialTimeout: 5 * time.Second,
	}

	c := make(chan os.Signal)
	signal.Notify(c)

	cli, err := clientv3.New(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	lockKey := "/lock"

	go func() {
		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}

		m := concurrency.NewMutex(session, lockKey)

		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go1 get mutex failed " + err.Error())
		}
		fmt.Printf("go1 get mutex sucess\n")
		fmt.Println(m)
		time.Sleep(time.Duration(10) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go1 release lock\n")
	}()

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go2 get mutex failed " + err.Error())
		}
		fmt.Printf("go2 get mutex sucess\n")
		fmt.Println(m)
		time.Sleep(time.Duration(2) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go2 release lock\n")
	}()

	<-c
}

func main2() {
	var conf = clientv3.Config{
		Endpoints:   []string{"localhost:23791", "localhost:23792", "localhost:23793"},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(conf)
	if err != nil {
		fmt.Print(err)
	}
	lease := clientv3.NewLease(client)
	leaseResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	leaseID := leaseResp.ID
	ctx, cancelFunc := context.WithCancel(context.TODO())

	//两个defer用于释放锁
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseID)

	//抢锁和占用期间，需要不停的续租，续租方法返回一个只读的channel
	keepChan, err := lease.KeepAlive(ctx, leaseID)
	if err != nil {
		fmt.Println(err)
	}
	//处理续租返回的信息
	go func() {
		for {
			select {
			case keepResp := <-keepChan:
				if keepChan == nil {
					fmt.Println("lease out")
					goto END
				} else {
					fmt.Println("get resp", keepResp.ID)
				}
			}
		}
	END:
	}()
	kv := clientv3.NewKV(client)
	txn := kv.Txn(context.TODO())
	//开始抢锁事务操作
	txn.
		If(clientv3.Compare(clientv3.CreateRevision("/lock/9"), "=", 0)).
		Then(clientv3.OpPut("/lock/9", "", clientv3.WithLease(clientv3.LeaseID(leaseID)))).
		Else(clientv3.OpGet("/lock/9"))
	//提交事务
	txnResp, err := txn.Commit()

	if err != nil {
		fmt.Println(err)
		return
	}
	//如果抢锁成功
	if txnResp.Succeeded {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
	time.Sleep(5 * time.Second)

}
