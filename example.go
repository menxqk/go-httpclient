package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/menxqk/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "BEarer ABC-123")

	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	createUser(User{"John", "Smith"})
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(bytes))
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(bytes))
}
