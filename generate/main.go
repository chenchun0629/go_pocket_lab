package main

import "fmt"

type Status int

//go:generate stringer -type Status,Color
const (
	Offline Status = iota
	Online
	Disable
	Deleted
)

type Color int

const (
	Write Color = iota
	Red
	Blue
)

func main() {
	fmt.Println(Write.String())
}
