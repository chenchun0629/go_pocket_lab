package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int
}

func main() {
	var v atomic.Value
	v.Store(&Config{})

	// 启动一个writer goroutine，不断写入数据
	go func() {
		i := 0

		for {
			i++
			cfg := &Config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	var ctx = context.TODO()
	ctx = context.WithValue(ctx, "key", "value")

	// 启动多个reader goroutine，不断获取数据
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load().(*Config)
				fmt.Printf("%#v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
