package main

import "fmt"

func main() {

	// basic loop in golang
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// for loop with 2 statements
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	fmt.Println()

	// other way to use for loop
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}
	fmt.Println()

	// breaking out of an infinite for loop
	b := 0
	for {
		fmt.Println(b)
		b++
		if b == 5 {
			break
		}
	}
	fmt.Println()

	// nested loops if golang
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}
	fmt.Println()

	// leaving out of a nested loop after a certain condition
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	fmt.Println()

	// using for loop to iterate over collections
	s := []int{3, 4, 5, 6}
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println()
}
