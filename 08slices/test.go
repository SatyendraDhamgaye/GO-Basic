package main

import (
	"fmt"
)

func main() {

	var hm = make(map[int]int)
	hm[1] = 10

	var nums = []int{1, 2, 3, 4}

	valuesss, ok := hm[10]
	fmt.Println(nums[2])
	fmt.Println(ok, valuesss)
}
