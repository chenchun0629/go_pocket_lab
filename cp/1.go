package main

import "fmt"

type s struct {
	a int
}

const INT_A = 1

func main() {
	var m = map[string]int{
		"a": 1,
	}
	fmt.Printf("%+v, %p \n", m, &m)

	var sa = s{
		a: 1,
	}
	var sb = sa
	sb.a = 2
	fmt.Printf("%+v, %+v \n", sa, sb)

	var spa = &s{
		a: 1,
	}
	var spc = *spa
	var spb = &spc
	spb.a = 2
	fmt.Printf("%+v, %+v \n", spa, spb)
	fmt.Printf("%p, %p, %p \n", spa, spb, &spc)
	fmt.Printf("%p, %p \n", &spa, &spb)

	var sla = []int{1, 2, 3}
	var slb = sla
	slb[0] = 0
	fmt.Printf("%+v, %+v \n", sla, slb)

	var sca = []int{1, 2, 3}
	var scb = make([]int, 3)
	copy(scb, sca)
	scb[0] = 0
	fmt.Printf("%+v, %+v \n", sca, scb)

	var ssa = []*s{{1}, {2}, {3}}
	var ssb = make([]*s, 3)
	copy(ssb, ssa)
	ssb[0].a = 0
	fmt.Printf("%+v, %+v \n", ssa, ssb)
	fmt.Printf("%+v, %+v \n", ssa[0].a, ssb[0].a)

}
