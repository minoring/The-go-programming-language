package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// defer schedules a function call to be run after the function completes.
// panic complete func immediately.
// recover stops the panic and returns the value that was passed to the call to panic. 
