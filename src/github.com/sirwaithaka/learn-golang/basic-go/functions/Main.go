package main

import "fmt"

// basic function in golang
func sayMessage() {
	fmt.Println("Hello Go!")
}

// basic function that takes parameters
func parameter(msg string) {
	fmt.Println(msg)
}

// function with multiple parameters
func multiple(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}

// a function with parameters of the same type
// can be written with the type at the end of
// the declarations of the parameters
func inferredType(greeting, name string) {
	fmt.Println(greeting, name)
}

// you can pass in pointers as parameters
func pointerParameter(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Tedd" // changes the value underlying the pointer
	fmt.Println(*name)
}

// variatic parameters in golang: take all the last arguments
// that are passed in and wrap them up in a slice that has the
// name of the variable in the last argument of the function
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// functions that return a result
func returnFunc(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// in golang functions can return a local variable as a pointer
// the go runtime will know that the return value is a pointer
// and the value will be moved automatically to the shared
// memory (heap) for us by the runtime.
func returnPointer(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// in golang return values can be named, in the function signature
// the return type can be named and the name will be made available
// in the local stack. Then that value will be implicitly returned.
func returnNamed(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return // implicit return of result variable
}

// in golang functions can return multiple values at the same time
// this is usually helpful with functions that could return errors
// or a computed value
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// anonymous functions in go
func anonymous() {
	for i := 0; i < 5; i++ {
		// define the anonymous function
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// functions declared as variables
func variableFuncs() {
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(6.9, 9.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

type greeter struct {
	greeting string
	name     string
}

// golang has the concept of methods that are abit different
// than the normal functions in go. A method is a function that
// is executing in a known context, a context is any type.

// define a method on the greeter type
// however this method get a copy of
// the greeter object so changing properties
// within this greeter will not reflect
// on the underlying object
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "" // this change on the greeter object will not reflect
}

// passing a pointer receiver to the method will make the
// method able to change values on the underlying object
func (s *greeter) salamu() {
	(*s).name = "Abdi"
	fmt.Println(s.greeting, s.name)
}

func methods() {

	g := greeter{
		greeting: "Hallooo",
		name:     "Golang!",
	}
	g.greet()
	g.salamu()

}

func main() {

	sayMessage()
	fmt.Println()

	parameter("hellooooo Go!")
	fmt.Println()

	for i := 0; i < 5; i++ {
		multiple("Hello too!", i)
	}
	fmt.Println()

	inferredType("Hellooo", "Stacey!")
	fmt.Println()

	greeting := "Hello"
	name := "Stanley"
	pointerParameter(&greeting, &name)
	fmt.Println(name)
	fmt.Println()

	sum("This is the sum ", 1, 2, 3, 4, 5, 6)
	fmt.Println()

	s := returnFunc(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Returned result ", s)
	fmt.Println()

	p := returnPointer(9, 8, 7, 6, 4, 8)
	fmt.Println("REturned value of pointer is ", *p)
	println()

	n := returnNamed(10, 39, 90)
	fmt.Println("Named return ", n)
	fmt.Println()

	_, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	dd, _ := divide(5.0, 2.0)
	fmt.Println(dd)
	fmt.Println()

	anonymous()
	fmt.Println()

	variableFuncs()
	fmt.Println()

	methods()
	fmt.Println()
}
