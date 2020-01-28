package main

import (
	"fmt"
	"log"
	"net/http"
)

// defer, panic, recover

func example1() {
	fmt.Println("start")
	// the defere keyword is eagerly evaluated here
	// but the execution of the statement is not until
	// the very end of the function
	defer fmt.Println("middle")
	fmt.Println("end")
}

func example2() {
	// here the defer statements are evaluated in last in
	// first out order. The last defer statement will be
	// executed first
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func example3() {
	// this example shows that variables passed
	// to deferred statements are eagerly evaluated
	// so even if the Print statement is executed last,
	// the value passed on to it is eagerly evaluated
	// and hence it will print "start"
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func example4() {
	// in this example, we use the panic function to
	// indicate a fatal exception that an application
	// cannot continue to execute following an error
	// example opening a web server to a port already
	// in use.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Go!"))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err.Error())
	}
}

func example5() {
	// this example demonstrate the sequence of execution
	// of deferred statements and panic statements
	// here the deferred is executed before the panic statement
	// is executed. Panic statements are executed after deferred
	// statements. Any statements in the function after the
	// panic statement are not going to be executed
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end") // this will not be executed
}

func example6() {
	// in this example we demonstrate the use of recover
	// function to catch a panicking application
	// The use of the recover function in the defer
	// statement will return nil if the function is not
	// panicking. If the function is panicking then the
	// recover function is going to return an error value
	// that is causing the application to panic
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Something has happened!")
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Something terrible has happened")
	fmt.Println("done panicking")
}

func example7() {
	// in this example we show how recover can be used
	// to catch panics in a deeper call stack. The
	// recover function ...

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func main() {

	example1()
	fmt.Println()

	example2()
	fmt.Println()

	// example of when to close a resource
	request()
	fmt.Println()

	example3()
	fmt.Println()

	// example4()
	// fmt.Println()

	// example5()
	// fmt.Println()

	example6()
	fmt.Println()

	example7()
	fmt.Println()

}
