package main

import (
	"fmt"
)

const (
	_ = iota
	kb = 1 << (iota * 10)	// kb = 1024
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

/*
        1024			10000000000
        1048576		100000000000000000000
        1073741824		1000000000000000000000000000000
*/
