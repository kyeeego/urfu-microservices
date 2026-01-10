package clients

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClientWithRetry interface {
	Get(url string, headers map[string]string) (int, []byte, error)
	Post(url string, headers map[string]string, body io.Reader) (int, []byte, error)
}

func NewHttpClientWithRetry(client *http.Client, attempts int, wait time.Duration) HttpClientWithRetry {
	return &endpointHttpClientImpl{
		client, attempts, wait,
	}
}

type endpointHttpClientImpl struct {
	client   *http.Client
	attempts int
	wait     time.Duration
}

func (c *endpointHttpClientImpl) Get(url string, headers map[string]string) (int, []byte, error) {
	var status int
	for range c.attempts {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("Unable to create request. Trying again: %e", err))
			time.Sleep(c.wait)
			continue
		}
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		response, err := c.client.Do(req)
		if err != nil {
			fmt.Println(fmt.Errorf("Error while sending request. Trying again: %e", err))
			time.Sleep(c.wait)
			continue
		}
		defer response.Body.Close()
		status = response.StatusCode

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(fmt.Errorf("Error while reading response. Trying again: %e", err))
			time.Sleep(c.wait)
			continue
		}

		return status, body, nil
	}

	return status, nil, fmt.Errorf("Unable to get response after %d attempts", c.attempts)
}

func (c *endpointHttpClientImpl) Post(url string, headers map[string]string, body io.Reader) (int, []byte, error) {
	var status int

	for range c.attempts {
		req, err := http.NewRequest(http.MethodPost, url, body)
		if err != nil {
			fmt.Println(fmt.Errorf("Unable to create request. Trying again: %e", err))
			time.Sleep(c.wait)
			continue
		}
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		response, err := c.client.Do(req)
		if err != nil {
			fmt.Println(fmt.Errorf("Error while sending request. Trying again: %e", err))
			time.Sleep(c.wait)
			continue
		}
		defer response.Body.Close()
		status = response.StatusCode

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(fmt.Errorf("Error while reading response. Trying again: %e", err))
			time.Sleep(c.wait)
			continue
		}

		return status, body, nil
	}

	return status, nil, fmt.Errorf("Unable to get response after %d attempts", c.attempts)
}
