package examples

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/menxqk/go-httpclient/gohttp"
)

func TestMain(m *testing.M) {
	// Tell the HTTP library to mock any further requests from here.
	gohttp.StartMockServer()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initalization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Fatal("and error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initalization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Fatal("and error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initalization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Errorf("no error was expected and we got '%s'", err.Error())
		}

		if endpoints == nil {
			t.Fatal("endpoints were expected and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid error message")
		}
	})
}
