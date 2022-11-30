package gohttp

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (c *httpClient) Get(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, mergeHeaders(headers), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, mergeHeaders(headers), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, mergeHeaders(headers), body)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, mergeHeaders(headers), body)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, mergeHeaders(headers), nil)
}

func (c *httpClient) Options(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, mergeHeaders(headers), nil)
}
