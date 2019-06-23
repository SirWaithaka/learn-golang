package main

import (
	"fmt"
	"strconv"
)

func variables() {
	var i int
	i = 42

	k := 99

	fmt.Println(i)
	fmt.Printf("%v %T\n", k, k)
}

func castVariables() {
	var i = 88
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}

func main() {
	variables()
	castVariables()
}
