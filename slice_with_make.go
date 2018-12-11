package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	
	b := make([]int, 5)
	printSlice("b", b)
	
	c := b[1:3]
	c[1] = 3
	printSlice("c", c)
	
	d := c[1:]
	printSlice("d", d)
	
	d[0] = 1
	printSlice("b", b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// a len=5 cap=5 [0 0 0 0 0]
// b len=5 cap=5 [0 0 0 0 0]
// c len=2 cap=4 [0 3]
// d len=1 cap=3 [3]
// b len=5 cap=5 [0 0 1 0 0]
