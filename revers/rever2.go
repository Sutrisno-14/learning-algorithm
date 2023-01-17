package main

import "fmt"

func main() {
	/*
		Using for range
	*/

	var result string
	const data = "try it"
	for _, v := range data {
		result = string(v) + result
	}
	fmt.Println(data)
	fmt.Println(result)

}
