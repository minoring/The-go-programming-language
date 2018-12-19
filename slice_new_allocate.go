package main

import (
	"fmt"
)

func main() {
	x := make([]int, 6, 7)
	a(x)
	fmt.Println(x)
	
	x = make([]int, 5, 5)
	a(x)
	fmt.Println(x)
}

func a(sl []int)() {
	sl = append(sl, 1)
	sl[0] = 5
	fmt.Println(sl)
}

/*
        [5 0 0 0 0 0 1]
        [5 0 0 0 0 0] // func main, len of x doesn't change.
        [5 0 0 0 0 1]
        [0 0 0 0 0]   // out of capacity, new allocation.
*/
