package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is a file which is create by os package -saty"

	file, err := os.Create("./myFile.txt")

	if err != nil {
		panic(err)
	}

	lenght, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("lenght of the file is: ", lenght)
	defer file.Close()

	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data inside file created is: ", string(databyte))
}
