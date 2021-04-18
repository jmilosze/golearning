package main

import (
	"fmt"
	"math/bits"
)

func main() {
	type a struct {
		b int
		c *int
	}

	fmt.Println(1<<(bits.UintSize-1) - 1)

	zxc := make(map[int]a)

	cv := 2
	zxc[0] = a{1, &cv}
}
