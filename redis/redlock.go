package main

import (
	"context"
	"fmt"
	"github.com/amyangfei/redlock-go/v2/redlock"
	"log"
	"time"
)

func main() {
	var endpoints = []string{
		"tcp://127.0.0.1:63801",
		"tcp://127.0.0.1:63802",
		"tcp://127.0.0.1:63803",
		"tcp://127.0.0.1:63804",
		"tcp://127.0.0.1:63805",
	}
	lockMgr1, err := redlock.NewRedLock(endpoints)
	if err != nil {
		log.Fatal(err, "====1")
	}

	lockMgr2, err := redlock.NewRedLock(endpoints)
	if err != nil {
		log.Fatal(err, "====2")
	}

	go func() {
		ctx := context.Background()
		if _, err := lockMgr1.Lock(ctx, "resource_name", 10*time.Second); err != nil {
			fmt.Println(err, "lockMgr1 lock err ====")
			return
		}
		fmt.Println("lockMgr1 lock ok")
	}()

	go func() {
		ctx := context.Background()
		if _, err := lockMgr2.Lock(ctx, "resource_name", 10*time.Second); err != nil {
			fmt.Println(err, "lockMgr2 lock err ====")
			return
		}

		fmt.Println("lockMgr2 lock ok")
	}()

	defer lockMgr1.UnLock(context.Background(), "resource_name")
	defer lockMgr2.UnLock(context.Background(), "resource_name")
	time.Sleep(time.Hour)
}
