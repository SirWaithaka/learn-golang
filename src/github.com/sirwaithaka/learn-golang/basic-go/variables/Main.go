package main

import (
	"fmt"
	"strconv"
)

func variables() {
	// declaring a variable
	var i int
	i = 42
	// var i int = 42

	k := 99

	fmt.Println(i)
	fmt.Printf("%v %T\n", k, k)

	var (
		firstName string = "John"
		lastName string = "Smith"
	)

	// code conventions
	// acronyms should be all uppercase
	var theURL string = "https://golang.org"
}

func castVariables() {
	var i = 88
	fmt.Printf("%v, %T\n", i, i)

	// converting int to string
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}

func main() {
	variables()
	castVariables()
}
