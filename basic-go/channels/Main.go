package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// an example of a channel that works with integers only
	ch := make(chan int)
	wg.Add(2)
	// goroutine that will be receiving
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

	// Now something important to keep in mind with channels is that default channels are
	// unbuffered and so when we try to send values to a channel and there is no receiver,
	// the application is going to panic.
	// wg.Add(1)
	// go func() {
	// 	// the receiving goroutine
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// for j := 0; j < 5; j++ {
	// 	wg.Add(1)
	// 	// We are going to spawn 5 goroutines writing to the channel
	// 	go func() {
	// 		ch <- 42
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// We can have goroutines that act as readers and writers at the same time
	wg.Add(2)
	go func() {
		i := <-ch      // wait for integer from channel
		fmt.Println(i) // print
		ch <- 1        // push value into channel
		wg.Done()
	}()
	go func() {
		ch <- 2           // push value into channel
		fmt.Println(<-ch) // wait for value and print it
		wg.Done()
	}()
	wg.Wait()

	// But a desirable functionality is having only one goroutine reading to a channel
	// and another writing to the channel. We can have this functionality by passing in
	// a channel as an argument to a goroutine with a bias to which direction we need
	// the goroutine to work on i.e. reading or writing.
	wg.Add(2)
	// Receving data from channel goroutine
	go func(ch <-chan int) {
		i := <-ch      // wait for integer from channel
		fmt.Println(i) // print
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 3 // push value into channel
		wg.Done()
	}(ch)
	wg.Wait()

	// Now we can have channels that have buffers, meaning that we can send multiple values
	// to a channel, and the channel will create an internal data store that will hold, these,
	// values until the processing functions can process them all.
	wg.Add(2)
	channel := make(chan int, 50)
	// Receving data from channel goroutine
	go func(ch <-chan int) {
		i := <-ch      // wait for integer from channel
		fmt.Println(i) // print
		wg.Done()
	}(channel)
	go func(ch chan<- int) {
		ch <- 4 // push value into channel
		ch <- 5 // this message is lost though
		ch <- 6 // this message is lost too
		wg.Done()
	}(channel)
	wg.Wait()

	// Now one way we can handle a case of reading multiple incoming values out of a channel
	// is by employing a looping construct on the channel
	chann := make(chan int, 50)
	wg.Add(2)
	// Receving data from channel goroutine
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(chann)
	go func(ch chan<- int) {
		ch <- 7
		ch <- 8
		close(ch) // then close the channel to let the looping construct know the chanel is closed
		wg.Done()
	}(chann)
	wg.Wait()

	// On the receiving side, we have some way to know if we are reading from a closed channel
	chann = make(chan int, 50)
	wg.Add(2)
	// Receving data from channel goroutine
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(chann)
	go func(ch chan<- int) {
		ch <- 9
		ch <- 10
		close(ch) // then close the channel to let the looping construct know the chanel is closed
		wg.Done()
	}(chann)
	wg.Wait()

	/* SELECT statements in channels */
	const (
		logInfo    = "INFO"
		logWarning = "WARNING"
		logError   = "ERROR"
	)

	type logEntry struct {
		time     time.Time
		severity string
		message  string
	}

	var logCh = make(chan logEntry, 50)
	var doneCh = make(chan struct{}) // This is a signal only channel

	logger := func() {
		for {
			select {
			case entry := <-logCh:
				fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
			case <-doneCh:
				break
			}
		}
	}

	logging := func() {
		go logger()
		logCh <- logEntry{time.Now(), logInfo, "App is starting"}

		logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
		time.Sleep(100 * time.Millisecond)
		doneCh <- struct{}{}
	}

	logging()
}
