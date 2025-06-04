package main

import (
	"io"
	"log"
	"strings"
)

func main() {
	payload := `{"user": "john", "action": "login"}`
	reader := strings.NewReader(payload)

	// read a rune
	r1, size, _ := reader.ReadRune()
	log.Printf("r1 %c", r1)   // {
	log.Println("size", size) // 1

	// read all
	content := make([]byte, reader.Len())
	_, err := reader.Read(content)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("content", string(content)) // content "user": "john", "action": "login"}

	// try to read again
	b, err := io.ReadAll(reader)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("content", string(b)) // content

	// try to read again after seeking
	// seek the reader
	reader.Seek(4, io.SeekStart)
	b, err = io.ReadAll(reader)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("content", string(b)) // content er": "john", "action": "login"}
	// returns 0 after ReadAll
	log.Println("len", reader.Len()) // 0

	// use a negative offset
	_, err = reader.Seek(-2, io.SeekEnd)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("len", reader.Len()) // 2

	_, err = reader.Seek(-2, io.SeekStart)
	if err != nil {
		log.Println("err", err) // err strings.Reader.Seek: negative position
		return
	}
}
