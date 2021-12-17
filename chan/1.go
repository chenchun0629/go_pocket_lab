package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
)

func main() {
	//fmt.Println(a1(10))

	a2()
}

func a2() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}

func fibonacci_chan(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
		fmt.Println(x)
		time.Sleep(time.Second)
	}
}

func a1(t int) error {
	errCh := make(chan error, t)
	wg := sync.WaitGroup{}
	wg.Add(t)
	for i := 0; i < t; i++ {
		go func(i int) {
			defer wg.Done()
			if i == 5 {
				errCh <- errors.New("hello cc")
			}
			fmt.Println("====", i)
			time.Sleep(time.Microsecond)
		}(i)
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
	//wg.Wait()
	//return nil
}
