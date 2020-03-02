package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// spawn a goroutine
	go sayHello()
	time.Sleep(10 * time.Millisecond)

	// We can use goroutines with closures as a concept
	// We can use variables defined in the outer scope of a
	// goroutine function. Example
	var msg = "Hello Closure"
	go func() {
		fmt.Println(msg) // prints "Hello Closure"
	}()
	time.Sleep(10 * time.Millisecond)
	// One thing to watch out in goroutines is a condition described as a
	// race condition. In the above closure example, anonymous the goroutine
	// is accessing a variable {msg} defined in the main function. But the goroutine
	// will not read the variable {msg} until the main function blocks. So this means
	// if the variable {msg} is changed before the goroutine runs, it would lead to
	// unexpected results. Example
	msg = "original message expected" // we expect this to be the message in the goroutine
	go func() {
		fmt.Println(msg) // prints "Changed message"
	}()
	msg = "Changed message"           // change content of variable, this will be printed instead
	time.Sleep(10 * time.Millisecond) // block the main function

	// We can go around the above race condition by passing the msg to the goroutine
	// instead
	msg = "passed message" // now the goroutine will print this
	go func(msg string) {
		fmt.Println(msg) // prints "passed message"
	}(msg)
	msg = "Change the message" // even if we change the variable {msg} later
	time.Sleep(10 * time.Millisecond)

	// Now for the next item, we really shouldnt be using time.Sleep to block the main
	// function so that we can allow our goroutines a chance to run. This is not a very
	// scalable solution. So Go programs usually employ a WaitGroup. Example of a WaitGroup
	// is declared above the main function.
	msg = "using wait group"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()

	// Now the other thing with WaitGroups is that they help you pause the execution of the main
	// program to give a chance to goroutines to execute too. But they done really order your goroutines
	// in an expected manner as to avoid race condtions. Example
	// In this example we expect the following value
	// Hello 0 Hello 1 Hello 2 Hello 3 Hello 4 Hello 5 Hello 6 Hello 7 Hello 8 Hello 9
	// But we get something like this
	// Hello #0 Hello #1 Hello #1 Hello #2 Hello #3 Hello #6 Hello #0 Hello #7 Hello #8 Hello #9
	var counter = 0
	sayhello := func() {
		fmt.Printf("Hello #%v ", counter)
		wg.Done()
	}
	increment := func() {
		counter++
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)      // We tell the wait group to wait on 2 goroutines
		go sayhello()  // prints "Hello {i}"
		go increment() // increments value of {i}
	}
	wg.Wait()
	fmt.Println()

	// So to tame the goroutines we can use something called a Mutex. It is a lock that the application
	// is going to honour. So basically if a mutex is locked on a value and a goroutine wants to update
	// the value, it will have to wait until it is unlocked.
	// So the following example has a result that looks like the following
	// Hello #1 Hello #1 Hello #1 Hello #2 Hello #2 Hello #2 Hello #2 Hello #2 Hello #2 Hello #3
	var m = sync.RWMutex{} // This is a Read-Write lock
	counter = 0
	sayhelloLock := func() {
		m.RLock() // Acquire a read lock
		fmt.Printf("Hello #%v ", counter)
		m.RUnlock() // release the lock
		wg.Done()
	}
	incrementLock := func() {
		m.Lock() // acquire a write lock
		counter++
		m.Unlock() // release the lock
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)          // We tell the wait group to wait on 2 goroutines
		go sayhelloLock()  // prints "Hello {i}"
		go incrementLock() // increments value of {i}
	}
	wg.Wait()
	fmt.Println()

	// The above example tries to tame the goroutines to follow some order which is good
	// but it still doesnt print the expected result. We can add a small change to how the locks
	// are acquired which should make it work.
	counter = 0
	sayhelloLock1 := func() {
		fmt.Printf("Hello #%v ", counter)
		m.RUnlock() // release the lock
		wg.Done()
	}
	incrementLock1 := func() {
		counter++
		m.Unlock() // release the lock
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)           // We tell the wait group to wait on 2 goroutines
		m.RLock()           // Acquire a read lock
		go sayhelloLock1()  // prints "Hello {i}"
		m.Lock()            // Acquire write lock
		go incrementLock1() // increments value of {i}
	}
	wg.Wait()
	fmt.Println()
}

func sayHello() {
	fmt.Println("hello")
}
