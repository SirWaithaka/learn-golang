package main

import (
	"fmt"
)

func main() {

	// declaring arrays
	// grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Println()

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "ahmed"
	students[2] = "arnold"
	fmt.Printf("Students: %v\n", students)

	// array of arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("identity matrix: %v\n", identityMatrix)

	// slices
	a := []int{1, 2, 3}
	b := a   // b will reference a
	b[1] = 5 // change in b reflects change in a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Length: %v\n", cap(a))
	fmt.Println()

	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Slice: %v\n", m[:])
	fmt.Printf("Slice: %v\n", m[3:])
	fmt.Printf("Slice: %v\n", m[:6])
	fmt.Printf("Slice: %v\n", m[3:6])
	fmt.Println()

	// using make function to create slices
	n := make([]int, 3, 100)
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	fmt.Println()

	// appending to slice
	p := []int{}
	fmt.Printf("Without elements: %v\n", p)
	fmt.Printf("Length: %v\n", len(p))
	fmt.Printf("Capacity: %v\n", cap(p))
	// append 1 value
	p = append(p, 1)
	fmt.Printf("Appended one element: %v\n", p)
	fmt.Printf("Length: %v\n", len(p))
	fmt.Printf("Capacity: %v\n", cap(p))
	// appending and array
	p = append(p, []int{2, 3, 4, 5}...)
	fmt.Printf("Appended array: %v\n", p)
	fmt.Printf("Length: %v\n", len(p))
	fmt.Printf("Capacity: %v\n", cap(p))
	fmt.Println()

	// Removing elements in a slice
	fmt.Printf("Slice: %v\n", p)
	fmt.Printf("Remove first element: %v\n", p[1:])
	fmt.Printf("Remove last element: %v\n", p[:len(p)-1])
	s := append(p[:2], p[3:]...)
	fmt.Printf("Remove 3rd element from slice: %v\n", s)
	fmt.Printf("Slice: %v\n", p)
	fmt.Println()
}
