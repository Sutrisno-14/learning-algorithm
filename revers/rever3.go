package main

import "fmt"

func main() {
	/*
		Using Function
	*/

	data := revers("try it")

	fmt.Println(data)
}

func revers(data string) string {
	var result string

	for i := 0; i < len(data); i++ {
		result += string(data[len(data) - 1 - i])
	}

	fmt.Println(data)
	return result
}
