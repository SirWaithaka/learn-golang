package requests

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	client *http.Client
}

type Header struct {
	key string
	values []string
}

type Request struct {
	method string
	url string
	headers []Header
	body io.Reader
}

type Body struct {
	io.ReadCloser
	json *json.Decoder
}

type Status struct {
	Code int
	Reason string
}

type Response struct {
	*Request
	Status Status
	Headers []Header
	Body
}

type ClientConfigOption func(*Client)

func NewClient(options ...ClientConfigOption) *Client {
	client := &Client{ client: http.DefaultClient }

	for _, opt := range options {
		opt(client)
	}
	return client
}

func WithTimeout(timeout time.Duration) ClientConfigOption {
	return func (c *Client) {
		c.client.Timeout = timeout
	}
}

func (c Client) IsSuccess(code int) bool {
	return code/100 == http.StatusOK/100
}

func (c *Client) Get(url string, options ...func(*Request)) (*Response, error) {
	request := &Request{
		method: "GET",
		url: url,
	}
	c.applyOptions(request, options...)

	return c.do(request)
}

func (c *Client) Post(url string, body io.Reader, options ...func(*Request)) (*Response, error) {
	request := &Request{
		method: "POST",
		url: url,
		body: body,
	}
	c.applyOptions(request, options...)

	return c.do(request)
}

func (c *Client) applyOptions(request *Request, options ...func(*Request)) {
	for _, opt := range options {
		opt(request)
	}
}

func (c *Client) do(request *Request) (*Response, error) {

	// check if Client.client is nil
	if c.client == nil {
		c.client = http.DefaultClient
		timeout := time.Duration(10 * time.Second)
		c.client.Timeout = timeout
	}

	req, err := c.newRequest(request)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, HTTPClientError{Message: err.Error()}
	}

	response := Response{
		Request: request,
		Status:  Status{ Code: res.StatusCode, Reason: res.Status[4:]},
		Headers: headers(res.Header),
		Body:    Body{ ReadCloser: res.Body },
	}
	return &response, nil
}

func (c *Client) newRequest(request *Request) (*http.Request, error) {
	req, err := http.NewRequest(request.method, request.url, request.body)
	if err != nil {
		return nil, HTTPClientError{Message: "could not build request object"}
	}

	req.Header = buildHeaders(request.headers)
	return req, nil
}

// IsJSON reads the returned headers and looks
// for the occurrence of 'Content-Type': 'application/json'.
// Returns true if the header exists, false otherwise.
func (r *Response) IsJSON() bool {
	headers := buildHeaders(r.Headers)

	var contains = func(search []string, key string) bool {
		str := strings.Join(search, ",")
		return strings.Contains(str, key)
	}

	return contains(headers["Content-Type"], "application/json")
}

// create map of header from slice of Header
func buildHeaders(headers []Header) map[string][]string {
	if len(headers) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for _, h := range headers {
		m[h.key] = h.values
	}
	return m
}

// create slice of Header from a map
func headers(h map[string][]string) []Header {
	headers := make([]Header, 0, len(h))
	for k,v := range h {
		headers = append(headers, Header{ key: k, values: v})
	}
	return headers
}

func WithHeader(key, value string) func(*Request) {
	return func(r *Request) {
		r.headers = append(r.headers, Header{
			key:key, values: []string{value},
		})
	}
}


//// parseUrl validates a given url
//func parseURL(url string) error {
//	u, err := nurl.Parse(url)
//	if err != nil {
//		return &InvalidURLError{url, "Is not a valid url"}
//	} else if u.Scheme == "" || u.Host == "" {
//		return &InvalidURLError{url, "Must be an absolute url"}
//	} else if u.Scheme != "http" && u.Scheme != "https" {
//		return &InvalidURLError{url, "Must begin with http or https"}
//	}
//
//	return nil
//}

