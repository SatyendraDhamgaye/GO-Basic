package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["ru"] = "RUBY"

	fmt.Println(languages)
	fmt.Println("js short for: ", languages["js"])
	fmt.Println("py short for: ", languages["py"])
	fmt.Println("ru short for: ", languages["ru"])

	delete(languages, "py")
	fmt.Println("After deleteing Python from the languages map: ", languages)

	//to check if the value is present in hash map or not
	//where value is value
	//ok return bool value i.e. true or false
	value, ok := languages["js"]
	fmt.Println(ok, value)

}
