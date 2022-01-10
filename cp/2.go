package main

import (
	"fmt"
	"pocket_lab/t_init"
	"reflect"
	"time"
)

func init() {
	fmt.Println("2")
}

func main() {
	t_init.A()
	var a = []int{1, 2, 3}
	var b = []int{1, 2, 3}
	//fmt.Printf("%v \n", a == b)
	fmt.Printf("%v \n", reflect.DeepEqual(a, b))

	var ch = cha()
	fmt.Println(<-ch)

	time.Sleep(time.Hour)

}

func cha() chan int {
	var ch = make(chan int, 1)

	time.AfterFunc(time.Second*3, func() {
		ch <- 1
	})

	return ch
}
