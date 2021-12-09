package main

import (
	"fmt"
	"hash/crc32"
	"stathat.com/c/consistent"
)

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func String2(s string) int {
	var (
		//l = len(s)
		r = 0
	)
	for _, v := range s {
		r += 31*r + int(v)
	}

	return r
}

func main() {

	var c = consistent.New()
	c.Add("haha")
	c.Add("hehe")
	c.Add("xixi")
	fmt.Println(c.Get("hoho"))
	//var s = "ccc"
	//var m = murmur3.New128()
	//_, _ = m.Write([]byte(s))
	//fmt.Println(string(m.Sum([]byte("aaa"))))
	//fmt.Println(String2(s), String(s), string(murmur3.New128().Sum([]byte(s))))
}
