package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Using POSTFORM method from the server")
	postFormReq()

}

func postFormReq() {

	myUrl := "http://localhost:8080/postform"

	data := url.Values{}
	data.Add("name", "satye")
	data.Add("email", "satye@gmail.com")

	response, err := http.PostForm(myUrl, data)
	catchError(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	var responseBuilder strings.Builder
	responseBuilder.WriteString(string(content))

	fmt.Println(responseBuilder.String())
}

func catchError(err error) {
	if err != nil {
		panic(err)
	}
}
