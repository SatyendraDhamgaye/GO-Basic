package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Using POST method from the server")
	postJsonRequest()

}

func postJsonRequest() {
	const myUrl = "http://localhost:8080/post"

	//json data to be send

	requestBody := strings.NewReader(`
	{
		"name" : "zero!!",
		"email" : "Dumb"
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	catchError(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	catchError(err)

	var responseBuilder strings.Builder
	responseBuilder.Write(content)
	fmt.Println(responseBuilder.String())
}

func catchError(err error) {
	if err != nil {
		panic(err)
	}
}
