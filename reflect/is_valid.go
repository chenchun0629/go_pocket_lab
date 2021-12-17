package main

import (
	"fmt"
	"reflect"
)

func main() {
	type s struct {
		name string
	}
	var x *s
	var rv = reflect.ValueOf(x)

	var i *int
	var v = reflect.ValueOf(i)
	fmt.Println(rv.Elem().IsValid())
	fmt.Println(v.Elem().IsValid())
	fmt.Println(reflect.Value{}.IsValid())
	fmt.Println(reflect.ValueOf(nil).IsValid())
}
