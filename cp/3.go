package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var h = Constructor(5, []int{5, 3, 6, 8, 9, 0, 1})
	h.Add(10)
	h.Add(11)
	h.Add(12)
	h.Add(13)
	h.Add(14)
	h.Add(15)
	fmt.Println(h.IntSlice)
}

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}
