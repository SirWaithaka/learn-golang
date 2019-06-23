package main

import (
	"fmt"
)

func main() {
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	k := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", m, m)

	var u uint16 = 65
	fmt.Printf("%v, %T\n", u, u)

	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	s := "this is a string"
	b := []byte(s)
	fmt.Printf("$v, %T\n", b, b)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
