package main

import "fmt"

type T struct {
	F1 string
	F2 []int
}

func main() {
	p := T{"some string", []int{10, 20}}
	q := p

	q.F1 = "zxc"
	q.F2[0] = 100

	fmt.Println(p)
	fmt.Println(q)
}
