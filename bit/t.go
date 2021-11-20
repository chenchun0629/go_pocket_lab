package main

import "fmt"

func main() {

	fmt.Println(abs(-1))

	//base()

	//xor()
	//and()
	//var a = 0b1100
	//	fmt.Printf(" %b \n", a & (a - 1))

	var a = 9
	fmt.Println(hasOneCount(a))
}

func hasOneCount(a int) int {
	var count = 0
	for a != 0 {
		a = a & (a - 1)
		count++
	}

	return count
}

func base() {
	fmt.Printf("与运算: %b \n", 0b10011&0b11001)
	fmt.Printf("或运算: %b \n", 0b10011|0b11001)
	fmt.Printf("异或运算: %b \n", 0b10011^0b11001)
	fmt.Printf("取反运算: %b \n", ^0b10011)
	fmt.Printf("左移运算: %d \n", 8<<3)
	fmt.Printf("右移运算: %d \n", 8>>3)
}

func and() {
	fmt.Println(1 & 1)
	fmt.Println(2 & 1)
	fmt.Println(3 & 1)
	fmt.Println(4 & 1)
}

func xor() {
	fmt.Println(1001 ^ 1101 ^ 1001)
	var a, b = 7, 9
	a, b = swap(a, b)
	fmt.Println(a, b)

	fmt.Println(^-7 + 1)

	fmt.Println(abs(-3))
}

func reversal(a int) int {
	return ^a + 1
}

func is_odd(a int) bool {
	return 1 == (a & 1)
}

func swap(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func abs(a int) int {
	var i = a >> 31
	return (a ^ i) - i
	//if i == 0 {
	//	return a
	//} else {
	//	return ^a + 1
	//}
}
