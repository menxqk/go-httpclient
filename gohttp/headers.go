package gohttp

import (
	"net/http"

	"github.com/menxqk/go-httpclient/gomime"
)

func mergeHeaders(headers []http.Header) http.Header {
	result := http.Header{}
	for _, h := range headers {
		for k, v := range h {
			if len(v) > 0 {
				result.Set(k, v[0])
			}
		}
	}

	return result
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add common headers to the request
	for header, value := range c.builder.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Set User-Agent if it is defined and not there yet
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}

	return result
}
