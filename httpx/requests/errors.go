package requests

import "fmt"

type HTTPClientError struct {
	Message string
}

func(err HTTPClientError) Error() string {
	return err.Message
}

// InvalidURLError is error type returned when URL passed is not good.
type InvalidURLError struct {
	url     string
	message string
}

func (e *InvalidURLError) Error() string {
	return fmt.Sprintf("Could not validate url %v. %v", e.url, e.message)
}

