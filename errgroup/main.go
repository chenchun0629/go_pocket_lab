package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	var eg, ctx = errgroup.WithContext(context.Background())

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

	_ = ctx
	_ = eg.Wait()
}
