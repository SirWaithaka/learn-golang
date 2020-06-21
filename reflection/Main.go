package main

import (
	"fmt"
	"os"
	"reflect"
)

type student struct {
	Fname, Lname string
	age uint64
}

func (s student) FullName() string {
	return s.Fname + " " + s.Lname
}

func main() {
	// this is an introduction to the laws of reflection
	// we start with reflection on primitive types of go
	var number = 5.25
	var num float64 = number
	fmt.Printf("Value of num %v\n", reflect.ValueOf(num)) // prints the value stored in the variable num i.e. 5.25
	fmt.Printf("Type of num %v\n", reflect.TypeOf(num)) // prints the type of the variable num i.e. float64
	// Kind is a function that checks the go type of a variable i.e. (u)int(8,16,32,64) or float(32,64) or struct or string
	// so a variable initialized with `type MyInt int` e.g. var num MyInt = 402 will return the Kind int instead of MyInt
	fmt.Printf("Kind of num %v\n", reflect.TypeOf(num).Kind())
	fmt.Printf("Kind of num %v\n", reflect.ValueOf(num).Kind())


	// to modify a reflection object, the value must be settable. To make a value settable we use a pointer to the variable.
	v := reflect.ValueOf(&num)
	// we can check if a reflection object is settable first with the function CanSet()
	fmt.Printf("Can Set %v\n", v.CanSet()) // this will return false

	// elem
	elem := v.Elem()
	fmt.Printf("Elem fn returns %v\n", elem)
	// using pointer object reflection we can use Set fns on Elem() return values to change the underlying values
	fmt.Println("Using reflection objects set fns, we change the value of a variable")
	elem.SetFloat(7.875)
	fmt.Printf("Reflection object value %v\n", elem)
	fmt.Printf("Original value %v\n", number)

	stud := student{"Mary", "Beelof", 18}

	fmt.Printf("Value of stud %v\n", reflect.ValueOf(&stud)) // prints &{Mary Beelof 18}
	fmt.Printf("Type of stud %v\n", reflect.TypeOf(stud)) // prints main.student
	fmt.Printf("Kind of stud %v\n", reflect.TypeOf(stud).Kind()) // prints struct

	// The listing of the field names of a reflection object from struct variable.
	for i := 0; i < reflect.ValueOf(&stud).Elem().NumField(); i ++ {
		fmt.Printf("Fields of the student struct %v: %v\n", i, reflect.ValueOf(stud).Type().Field(i).Name)
	}

	// we can modify the field properties of a struct using reflection
	reflected := reflect.ValueOf(&stud).Elem()
	for i := 0; i < reflected.NumField(); i ++ {
		prop := reflected.Field(i)
		if prop.Kind() == reflect.String {
			prop.SetString(os.ExpandEnv(prop.String()))
		}
	}

	fmt.Printf("After mutation of stud %v\n", reflect.ValueOf(&stud))

}