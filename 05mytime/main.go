package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Time: ")
	time := time.Now()

	fmt.Println(time)
	fmt.Println(time.Format("01-02-2006 15:04:05 Monday"))
	//;(

}
