package main

import (
	"fmt"
)

func main() {
	//weeks := []string{"monday", "tuesday", "wednesday", "thrusday", "friday", "saturday", "sunday"}

	// for i := 0; i < len(weeks); i++ {
	// 	fmt.Println(weeks[i])
	// }

	// for i := range weeks {
	// 	fmt.Println(weeks[i])
	// }

	// for index, day := range weeks {
	// 	fmt.Printf("Index is %v and day is %v\n", index, day)
	// }

	val := 0

	for val < 10 {
		if val == 2 {
			goto out
		}
		if val >= 8 {
			break
		}

		if val == 5 {
			val++
			continue
		}
		fmt.Println(val)
		val++
	}

out:
	fmt.Println("Jumping out")
}
