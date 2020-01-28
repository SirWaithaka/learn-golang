package service

import (
	"fmt"
	"io"
	"net/http"
	nurl "net/url"
	"time"
)

var timeout time.Duration = time.Duration(5 * time.Second)

type HTTPClient struct {
	Client  http.Client
	BaseURL string
}

func NewClient() http.Client {
	// change the default timeout for the http client
	// TODO: build the client and inject it.
	return http.Client{
		Timeout: timeout,
	}
}

func NewHTTPClient(c http.Client, baseURL string) *HTTPClient {
	return &HTTPClient{c, baseURL}
}

func (c HTTPClient) IsSuccess(code int) bool {
	return code/100 == http.StatusOK/100
}

func (c *HTTPClient) Do(method, url string, body io.Reader) (*http.Response, error) {

	// append base url to passed in url
	u := c.BaseURL + url

	var err error
	// confirm url is valid
	err = parseURL(u)
	if err != nil {
		return nil, err
	}

	// prepare a request object
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request %v: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return c.Client.Do(req)
}

func (c *HTTPClient) Get(url string) (*http.Response, error) {
	return c.Do(http.MethodGet, url, nil)
}

func (c *HTTPClient) Post(url string, body io.Reader) (*http.Response, error) {
	return c.Do(http.MethodPost, url, body)
}

func parseURL(url string) error {
	u, err := nurl.Parse(url)
	if err != nil {
		return fmt.Errorf("error: %v is not a valid url", url)
	} else if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("error: %v must be an absolute url", url)
	} else if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("error: %v must begin with http or https")
	}

	return nil
}
