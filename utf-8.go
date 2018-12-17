package main

import (
	"fmt"
)

func main() {
	s := "Hello한日"     // the explicit UTF-8 bytes. 日 for 3bytes.
	fmt.Println(len(s)) // 11 bytes. 11 len.
	fmt.Printf("%T\n", s)
	
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
}

/*
        11
        string
        [72 101 108 108 111 237 149 156 230 151 165]
        []uint8
        U+0048 'H'
        U+0065 'e'
        U+006C 'l'
        U+006C 'l'
        U+006F 'o'
        U+00ED 'í'
        U+0095
        U+009C
        U+00E6 'æ'
        U+0097
        U+00A5 '¥'
*/
