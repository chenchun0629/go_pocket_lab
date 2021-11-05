package main

import "fmt"

func main() {
	var a = new(A)
	fmt.Println(a)
}

type A struct {
	S string
}
