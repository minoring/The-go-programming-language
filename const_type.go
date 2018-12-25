package main

import (
	"fmt"
)

type foo int

var y foo

const bar = 42 // What is type of bar? Not determinant.

func main() {
	y = bar
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)
}

/*
main.foo
int
42

Named types vs anonymous types

Anonymous types are indeterminate. They have not been declared as a type yet. 
The compiler has flexibility with anonymous types. 
You can assign an anonymous type to a variable declared as a certain type. 
If the assignment can occur, the compiler will figure it out; the compiler will do an implicit conversion. 
You cannot assign a named type to a different named type.

*/
