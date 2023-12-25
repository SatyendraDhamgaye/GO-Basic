package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://satyendradhamgaye52002.on.drv.tw:8000/www.Satyendra's_CV.com?data=satyendra&pay=999999"

func main() {
	//fmt.Println(url)

	result, err := url.Parse(myurl)
	errorCatcher(err)

	fmt.Println("Host: ", result.Host)
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	params := result.Query()

	fmt.Println("There are two key pairs in the url")
	fmt.Println("Key: data", params["data"])
	fmt.Println("Key: pay", params["pay"])

	for _, val := range params {
		fmt.Println(val)
	}

	//creating a url using various data

	createUrl := &url.URL{
		Scheme:  "https",
		Host:    "abc.com",
		Path:    "def",
		RawPath: "ghi",
	}

	fmt.Println("Created Url: ", createUrl.String())
}

func errorCatcher(err error) {
	if err != nil {
		panic(err)
	}
}
