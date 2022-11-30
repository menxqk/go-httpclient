package gohttp

import (
	"net/http"
	"time"
)

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder

	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

type clientBuilder struct {
	Headers             http.Header
	maxIndleConnections int
	connectionTimeout   time.Duration
	responseTimeout     time.Duration
	disableTimeouts     bool
	baseUrl             string
	client              *http.Client
	userAgent           string
}

func (cb *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	cb.Headers = headers
	return cb
}

func (cb *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	cb.connectionTimeout = timeout
	return cb
}

func (cb *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	cb.responseTimeout = timeout
	return cb
}

func (cb *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	cb.maxIndleConnections = connections
	return cb
}

func (cb *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	cb.disableTimeouts = disable
	return cb
}

func (cb *clientBuilder) Build() Client {
	client := &httpClient{
		builder: cb,
	}
	return client
}

func (cb *clientBuilder) SetHttpClient(c *http.Client) ClientBuilder {
	cb.client = c
	return cb
}

func (cb *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	cb.userAgent = userAgent
	return cb
}
