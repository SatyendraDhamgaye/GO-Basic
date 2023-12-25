package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://satyendradhamgaye52002.on.drv.tw/www.Satyendra's_CV.com/"

func main() {
	fmt.Println("Handeling web requests")

	response, err := http.Get(url)

	errorCatcher(err)

	fmt.Printf("Response is of Type: %T", response)

	defer response.Body.Close() //always close it, it's not done automatically

	dataByte, err := ioutil.ReadAll(response.Body)

	errorCatcher(err)

	fmt.Println("Data inside the given url is: ", string(dataByte))
}

func errorCatcher(err error) {
	if err != nil {
		panic(err)
	}
}
