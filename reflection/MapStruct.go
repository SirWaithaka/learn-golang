package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Header map[string][]string

type HeaderValues struct {
	MSISDN string `header:"User-msisdn"`
	IMEI   string `header:"User-imei"`
}

func readTags(headers Header, v interface{}) error {
	var tagName = "header"

	var t = reflect.TypeOf(v)
	var k = t.Kind()

	// we check if the parameter is a pointer
	if t.Kind() != reflect.Ptr {
		return errors.New("value should be a pointer")
	}

	valPtr := reflect.Indirect(reflect.ValueOf(v))
	log.Printf("Value of pointer %v", valPtr)

	t = valPtr.Type()
	k = valPtr.Kind()

	if k != reflect.Struct {
		return errors.New("type of value should be a struct")
	}

	log.Println("Working with a struct")

	reflected := reflect.ValueOf(v).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := reflected.Type().Field(i)
		prop := reflected.Field(i)

		tag := field.Tag.Get(tagName)
		h := headers[tag]
		prop.SetString("")
		if len(h) != 0 {
			prop.SetString(headers[tag][0])
		}


		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}

	return nil
}

func main() {
	h := Header{
		"User-Msisdn": []string{"254712345678"},
		"User-imei":   []string{"453256366747"},
	}

	number := 30

	var err error
	err = readTags(h, HeaderValues{})
	if err != nil {
		log.Printf(err.Error())
	}

	var he HeaderValues
	err = readTags(h, &he)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("New value of struct %+v", he)

	err = readTags(h, new(HeaderValues))
	if err != nil {
		log.Printf(err.Error())
	}

	err = readTags(h, new(interface{}))
	if err != nil {
		log.Printf(err.Error())
	}

	err = readTags(h, &number)
	if err != nil {
		log.Printf(err.Error())
	}

	value := reflect.ValueOf(new(HeaderValues)).Type()
	fmt.Printf("Value of a pointer %v", value)
}
