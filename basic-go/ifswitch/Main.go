package main

import (
	"fmt"
)

func main() {

	number := 50
	guess := -5

	// guess is less than 1 or greater than 100
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	//  example of a switch statement
	switch 2 {
	case 1:
		fmt.Println("matched one")
	case 2:
		fmt.Println("matched two")
	default:
		fmt.Println("no matched value")
	}

	// switch with multiple options in a case
	switch 2 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 6, 8:
		fmt.Println("two, 6 or eight")
	default:
		fmt.Println("no matched value")
	}

	// swtich case with conditional
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less thatn or equal to 10")
	case i <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than twenty")
	}

	// special use of a type switch
	var a interface{} = [3]int{}
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	case float64:
		fmt.Println("a is an int")
	case string:
		fmt.Println("a is an int")
	case [3]int:
		fmt.Println("a is an array")
	default:
		fmt.Println("a is another type")
	}

	// breaking from a test case early
	// use the keyword break to leave a test
	// case early
	var b interface{} = 4
	switch b.(type) {
	case int:
		fmt.Println("b is an int")
		break // leave case
		// this line will not execute
		fmt.Println("Yes b is an integer")
	case float64:
		fmt.Println("b is an int")
	case string:
		fmt.Println("b is an int")
	case [3]int:
		fmt.Println("b is an array")
	default:
		fmt.Println("b is another type")
	}
}
