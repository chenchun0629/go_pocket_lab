package main

import (
	"github.com/pkg/errors"
)

func main() {
	var err = errors.New("err")
	_ = err

	var ch = make(chan string)
	_ = ch
}
