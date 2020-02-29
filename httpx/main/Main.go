package main

import (
	//"io/ioutil"
	"log"
	"time"

	"requests"
)

func main() {

	c := requests.NewClient(requests.WithTimeout(time.Duration(10 * time.Second)))

	//url := "https://youtise-location-dev.herokuapp.com/api/location/gitari@somethingincltd.com"
	//response, err := c.Get(url, requests.WithHeader("Accept", "application/json"))
	response, err := c.Get("https://google.com", requests.WithHeader("Accept", "application/json"))
	//response, err := c.Get("https://httpbin.org/get", requests.WithHeader("Accept", "application/json"))
	if err != nil {
		log.Println(err)
	}

	//bytes, _ := ioutil.ReadAll(response.Body)

	log.Println(response.Status.Reason)
	//log.Println(response.Headers)
	//log.Println(string(bytes))

	log.Println(response.IsJSON())
}
