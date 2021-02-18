package main

import "fmt"

func main() {

	academy := "Alterra"

	var reverse string
	for _, val := range academy {
		reverse = string(val) + reverse
	}

	fmt.Println(academy)
}
