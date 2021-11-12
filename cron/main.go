package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	_, _ = c.AddFunc("@every 1s", func() {
		fmt.Println("begin tick every 1 second")
		time.Sleep(5 * time.Second)
		fmt.Println("end tick every 1 second")
	})

	c.Start()
	time.Sleep(time.Second * 11)
	fmt.Println("==============1")
	var ctx = c.Stop()
	fmt.Println("==============2")
	<-ctx.Done()
	fmt.Println("==============3")
}
