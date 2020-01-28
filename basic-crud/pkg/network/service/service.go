package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func NewApiService(c *HTTPClient) *ApiService {
	return &ApiService{c}
}

type ApiService struct {
	client *HTTPClient
}

type locationEndpoints struct {
	RouteAuthenticate string // the page where signup happens
	RouteAdverts      string // the page for getting adverts
	RouteLocations    string // the page for getting locations
	RouteAdvertLog    string // the page for posting advert logs
	RouteConversion   string // the page for recording conversions
	RouteDataPhoto    string // the page for posting data picture
}

type Routing struct {
	LocationApi locationEndpoints
}

var routes Routing = Routing{
	LocationApi: locationEndpoints{
		RouteAuthenticate: "location/login",
		RouteAdverts:      "location/adverts",
		RouteLocations:    "location",
		RouteAdvertLog:    "analytics/create-adlog",
		RouteConversion:   "analytics/conversion",
		RouteDataPhoto:    "analytics/record-ad-analytic",
	},
}

func (api ApiService) Authenticate(placeId, password string) ([]byte, error) {
	// create post data for the http client
	dataMap := map[string]string{
		"placeId":  placeId,
		"password": password,
	}
	postData, _ := json.Marshal(dataMap)
	reader := bytes.NewReader(postData)

	resp, err := api.client.Post(routes.LocationApi.RouteAuthenticate, reader)
	if err != nil {
		return nil, fmt.Errorf("error: could not perform http request %v", err)
	}
	defer resp.Body.Close()

	// convert reader into byte array
	byteBody, _ := ioutil.ReadAll(resp.Body)

	// return error if status code is not successful
	if !api.client.IsSuccess(resp.StatusCode) {
		return nil, fmt.Errorf("error: %v", string(byteBody))
	}

	return byteBody, nil
}
