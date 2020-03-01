# A Simple Expressive HTTP requests client
This here is a simple client I use to perform http requests. Implemented methods are Get, Post. Maybe one
day i will add the rest :-)
The client can be easily configured and extended, depending on the usecase.

## Installation

To run the example.
```bash
$ cd main
$ go build Main.go
$ ./Main
```

## Example Uses

```go

package main

import (
	"io/ioutil"
	"log"
	"time"

	"requests"
)

func main() {

	c := requests.NewClient(requests.WithTimeout(time.Duration(10 * time.Second)))
	
    //response, err := c.Get("https://httpbin.org/get", requests.WithHeader("Accept", "application/json"))
	response, err := c.Get("https://google.com", requests.WithHeader("Accept", "application/json"))
	if err != nil {
		log.Println(err)
		return
	}

	bytes, _ := ioutil.ReadAll(response.Body)

	log.Println(response.Status.Reason) // prints status string e.g. OK for 200
	log.Println(response.Headers) // print out the response headers
	log.Println(string(bytes)) // print out the results in string
    
	// Returns a bool indicating the response is json valid
	log.Println(response.IsJSON()) // this example returns false, uncomment httpbin url request to return true
}

```