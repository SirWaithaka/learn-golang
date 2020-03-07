package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"

	//"io/ioutil"
	"log"
	"time"

	"requests"
)

func main() {

	c := requests.NewClient(requests.WithTimeout(time.Duration(10 * time.Second)))

	//url := "https://youtise-location-dev.herokuapp.com/api/location/gitari@somethingincltd.com"
	//response, err := c.Get(url, requests.WithHeader("Accept", "application/json"))
	//url := "https://youtise-location-dev.herokuapp.com/api/location/adverts"
	//response, err := c.Get(url, requests.WithHeader("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhhMzlhZmQwLTcwM2EtNDAxZC1iMWY1LTIwZmI3ODYxNWQ4ZiIsInNlcmlhbE51bWJlciI6IjRDNEM0NTQ0LTAwMzUtMzIxMC04MDRCLUI0QzA0RjU4NTMzMiIsInBsYWNlSWQiOiI2MWY3MGIyMy0zYWMyLTQ4NzQtOTA4YS00MjU1YjA2YjNhZDMiLCJpYXQiOjE1ODI3MjAwNTgsImV4cCI6MTYxNjU4ODg1OCwiaXNzIjoiU29tZXRoaW5nSW5jIEx0ZCJ9.aByzusDBXJ71TQjSwnB3_qZvcRk8Z7sS_55rwT1Mark"))

	type PostParams map[string]interface{}

	data := PostParams {
		"id" : "61f70b23-3ac2-4874-908a-4255b06b3ad3",
		"password" : strconv.Itoa(1234),
		"serialNumber" : "4C4C4544-0035-3210-804B-B4C04F585332",
	}
	// convert map into byte array and later into Reader
	postData, _ := json.Marshal(data)
	log.Printf("marshalled %s", string(postData))
	reader := bytes.NewReader(postData)
	//url := "https://youtise-location-dev.herokuapp.com/api/location/login"
	url := "http://localhost:6000/api/location/login"
	response, err := c.Post(
		url,
		reader,
		requests.WithHeader("Accept", "application/json"),
		//requests.WithHeader("Content-Type", "application/x-www-form-urlencoded"),
		requests.WithHeader("Content-Type", "application/json"),
		)

	//response, err := c.Get("https://google.com", requests.WithHeader("Accept", "application/json"))
	//response, err := c.Get("https://httpbin.org/get", requests.WithHeader("Accept", "application/json"))
	if err != nil {
		log.Println(err)
	}

	byteSlice, _ := ioutil.ReadAll(response.Body)

	log.Println(response.Status.Reason)
	//log.Println(response.Headers)
	log.Println(string(byteSlice))

	log.Println(response.IsJSON())
}
