package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	var eg, ctx = errgroup.WithContext(context.Background())

	var cctx, cancle = context.WithCancel(ctx)

	eg.Go(func() error {
		<-cctx.Done()
		fmt.Println("cctx done ...")
		return nil
	})

	eg.Go(func() error {
		fmt.Println("====================1")
		time.Sleep(time.Second * 5)
		fmt.Println("====================2")
		return nil
	})

	eg.Go(func() error {
		fmt.Println("====================a")
		time.Sleep(time.Second * 3)
		fmt.Println("====================b")
		return nil
	})
	fmt.Println("=========bc")
	cancle()
	fmt.Println("=========ec")

	_ = eg.Wait()
}
