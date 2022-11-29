package gohttp

import (
	"encoding/xml"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")

	client := NewBuilder().SetHeaders(commonHeaders).Build()

	c, ok := client.(*httpClient)
	if !ok {
		t.Fatalf("wrong concrete type for http client: %T", client)
	}

	requestHeader := make(http.Header)
	requestHeader.Set("X-Request-Id", "ABC-123")

	// Execution
	finalHeaders := c.getRequestHeaders(requestHeader)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("expected 3 headers{")
	}
	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}
	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}
}

func TestGetRequestBody(t *testing.T) {
	// Initialization
	client := NewBuilder().Build()

	c, ok := client.(*httpClient)
	if !ok {
		t.Fatalf("wrong concrete type for http client: %T", client)
	}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := c.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("no error expected when passing nil body")
		}
		if body != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"} // ["one","two"]
		body, err := c.getRequestBody("application/json", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		// Execution
		requestBody := struct {
			XMLName xml.Name `xml:"xml"`
			First   string   `xml:"first"`
			Second  string   `xml:"second"`
		}{
			First:  "three",
			Second: "four",
		}
		body, err := c.getRequestBody("application/xml", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling xml")
		}
		if string(body) != `<xml><first>three</first><second>four</second></xml>` {
			t.Error("invalid xml body obtained")
		}

	})
	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		// Execution
		requestBody := []string{"five", "six"} // ["one","two"]
		body, err := c.getRequestBody("", requestBody)

		// validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}
		if string(body) != `["five","six"]` {
			t.Error("invalid json body obtained")
		}

	})
}
