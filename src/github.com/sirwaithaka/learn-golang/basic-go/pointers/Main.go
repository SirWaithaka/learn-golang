package main

import "fmt"

func main() {

	var a int = 43
	var b *int // declaring a pointer to an integer
	b = &a     // b now holds the address of a in memory

	fmt.Println(a, b)

	// using dereferencing
	fmt.Println(*b) // pulls the value at location b

	*b = 10 // changing the value of *b changes the value of a
	fmt.Println(a, *b)

	// golang doesnt allow pointer arithmetic
	// e.g.
	// b + 4

	// pointer types in golang
	type aStruct struct {
		foo int
	}

	var ms *aStruct
	ms = &aStruct{foo: 42}
	fmt.Println(ms)

	// we can also use the new builtin fn to initialize pointers
	// to an object. But the new function doesnt allow initialization
	// of the object using the object initialization syntax. So the
	// object fields are going to be initialized to their zero values
	var as *aStruct
	fmt.Println(as)   // pointer that is not initialized is going to hold the value nil
	as = new(aStruct) // initializes to zero values
	fmt.Println(as)

	// accessing the fields in a struct pointer
	(*as).foo = 99
	fmt.Println((*as).foo)

}
