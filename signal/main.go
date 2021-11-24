package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP:
				fmt.Println("SIGHUP Program Exit...", s)
				GracefullExit()
			case syscall.SIGINT:
				fmt.Println("SIGINT Program Exit...", s)
				GracefullExit()
			case syscall.SIGTERM:
				fmt.Println("SIGTERM Program Exit...", s)
				GracefullExit()
			case syscall.SIGQUIT:
				fmt.Println("SIGQUIT Program Exit...", s)
				GracefullExit()
			//case syscall.SIGUSR1:
			//	fmt.Println("usr1 signal", s)
			//case syscall.SIGUSR2:
			//	fmt.Println("usr2 signal", s)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("Program Start...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func GracefullExit() {
	fmt.Println("Start Exit...")
	fmt.Println("Execute Clean...")
	fmt.Println("End Exit...")
	os.Exit(0)
}
