package qod

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const uri = "http://quotes.rest/"

type Client struct{}

func (c *Client) get(url string, headers map[string]string) (*http.Response, error) {
	req, err := buildRequest("GET", url, headers)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 400 {
		return nil, errors.New("400: bad request")
	}

	if res.StatusCode == 401 {
		return nil, errors.New("401: access denied")
	}

	if res.StatusCode == 403 {
		return nil, errors.New("403: forbidden")
	}

	if res.StatusCode == 404 {
		return nil, errors.New("404: not found")
	}

	return res, nil
}

func ReadAll(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildRequest(method, url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
