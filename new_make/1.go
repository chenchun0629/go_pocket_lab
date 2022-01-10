package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	var s *TestStruct
	fmt.Println(s == nil)    // #=> true
	fmt.Println(NilOrNot(s)) // #=> false

	var x = sync.Map{}

	go func() {
		for i := 0; i < 100; i++ {
			x.Store(fmt.Sprintf("x%d", i), i)
			fmt.Println("a====", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			_ = fmt.Sprintf("x%d", i)
			_ = fmt.Sprintf("x%d", i)
			_ = fmt.Sprintf("x%d", i)
			_ = fmt.Sprintf("x%d", i)
			fmt.Print("b====")
			fmt.Println(x.Load("x9"))
			//time.Sleep(time.Microsecond)
		}
	}()

	select {}
}

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	fmt.Println(reflect.ValueOf(v).IsNil())
	return v == nil
}

func main1() {
	var m = new(map[int]int)
	//(*m)[1] = 2 // panic
	fmt.Printf("%+v \n", m)

	var a *int
	fmt.Printf("%+v \n", a)
	a = new(int)
	fmt.Printf("%+v \n", a)
}
