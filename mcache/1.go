package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"fmt"
	"math/rand"
)

func main() {

	var num = getRandom()
	fmt.Printf("main num: %p \n", num)
	fmt.Println(*num)

	//var num2 = getRandom()
	//fmt.Printf("main num2: %p \n", &num2)

	go func() {
		var err = http.ListenAndServe("0.0.0.0:8081", nil)
		fmt.Println(err, "==========")
	}()

	time.Sleep(time.Hour)
}

func getRandom() *int {
	var tmp = rand.Intn(100)
	//fmt.Printf("getRandom tmp: %p \n", &tmp)
	return &tmp
}

//func main() {
//	var num = getRandom()
//	fmt.Printf("main num: %p \n", &num)
//	fmt.Println(num)
//}
//
func getRandom2() int {
	var tmp = rand.Intn(100)
	//fmt.Printf("getRandom2 tmp: %p \n", &tmp)
	return tmp
}
