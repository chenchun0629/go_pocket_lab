package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"time"
)

func main() {
	pool := tunny.NewFunc(10, func(payload interface{}) interface{} {
		fmt.Println(payload)
		time.Sleep(time.Second)
		return payload
	})
	defer pool.Close()

	for i := 0; i < 10; i++ {
		pool.Process(i)
	}
}
