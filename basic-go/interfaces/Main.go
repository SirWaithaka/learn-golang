package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
	fmt.Println()

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.increment())
	}
	fmt.Println()

	// using interfaces embedded in interface
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello youtube listeners, this is a test"))
	wc.Close()
	fmt.Println()

	// interface conversion -> converting a type interface into
	// another type that implements the same methods. In this case
	// the variable wc is an instance of WriterCloser but the instance
	// was created from the BufferedWriterCloser struct which has some
	// other implementation details in its properties. if we want to
	// access the buffer property of the wc variable we need to do
	// a conversion of the type into a BufferedWriterCloser
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)
	fmt.Println()

	// if we try to do a type conversion into a type that does not
	// implement the methods on an interface instance, the go runtime
	// is going to panic, however we can avoid this by using the ok syntax
	if r, ok := wc.(io.Reader); ok {
		fmt.Println(&r)
	} else {
		fmt.Println("conversion failed")
	}
}

// Interfaces describe relationships
// this interface describes an object that
// can given some bytes can write them to
// something.
type Writer interface {
	Write([]byte) (int, error)
}

// Normally structs are used to implement interfaces
// but any type can be used to implement an interaface
// we define here a struct ConsoleWriter, then we can
// define a method on the ConsoleWriter struct that
// implements the relationships of the Writer interface
type ConsoleWriter struct{}

// This is the method definition of the Write method of
// the Writer interface, that takes in the ConsoleWriter,
// therefore this is a method on the ConsoleWriter struct
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	increment() int
}

type IntCounter int

func (ic *IntCounter) increment() int {
	*ic++
	return int(*ic)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// write to the console in 8 character increments
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
