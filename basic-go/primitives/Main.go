package main

import (
	"fmt"
)

func main() {
	var n bool = false
	fmt.Println("The boolean value and type")
	fmt.Printf("%v, %T\n", n, n)
	fmt.Println()

	k := 1 == 1 // check lh value equals rh value
	m := 1 == 2 // check lh value equals rh value

	fmt.Println("Boolean checks")
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Println()

	// unsigned integers
	// uint8 [0-255]
	// uint16 [0-65535]
	// uint32 [0-4294967295]


	// byte is alias for uint8


	// arithmetic operations
	y := 10
	z := 3
	fmt.Println("Arithmentic Operations:")
	fmt.Printf("10 + 3: %v\n", (y + z))
	fmt.Printf("10 - 3: %v\n", (y - z))
	fmt.Printf("10 / 3: %v\n", (y / z))
	fmt.Printf("10 * 3: %v\n", (y * z))
	fmt.Printf("10 %% 3: %v\n", (y % z)) // remainder
	fmt.Println()


	// bitwise operations
	g := 10 // 1010
	h := 3 	// 0011
	fmt.Println("Bitwise Operations:")
	fmt.Printf("10 & 3: %v\n", (g & h)) // bitwise AND -> 0010
	fmt.Printf("10 | 3: %v\n", (g | h)) //bitwise OR -> 1011
	fmt.Printf("10 ^ 3: %v\n", (g ^ h))  //bitwise EXCLUSIVE-OR -> 1001
	fmt.Printf("10 &^ 3: %v\n", (g &^ h)) //bitwise AND-NOT -> 0100 
	fmt.Println()


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

	// zero values for all type
	// bool = false
	// int = 0
	// float32 = 0
}
