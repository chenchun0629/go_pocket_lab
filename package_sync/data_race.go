package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter int

func main() {
	for routine := 0; routine < 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("final count: %d \n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		Counter = Counter + 1
		//value := Counter
		//time.Sleep(time.Nanosecond)
		//value++
		//Counter = value
	}
	Wait.Done()
}
