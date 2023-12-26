package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Using GET method from the server")
	performGetReq()
}

func performGetReq() {
	const myUrl = "http://localhost:8080/get"

	response, err := http.Get(myUrl)
	catchError(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseBuilder strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	bytecount, _ := responseBuilder.Write(content)
	catchError(err)
	fmt.Println("Byte count from the `strings.Builder`: ", bytecount)
	fmt.Println(responseBuilder.String())
	//fmt.Println("Content of the Body is: ", string(content))

}

func catchError(err error) {
	if err != nil {
		panic(err)
	}
}
