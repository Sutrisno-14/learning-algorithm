package main

import "fmt"

func main() {
	// var data, result string

	// data = "KASUR"

	// for i:=0; i < len(data)/2; i++ {
	// 	dataReserv := len(data) - 1 - i

	// 	if data[i] == data[dataReserv] {
	// 		result = "it is palindrome"
	// 	}else {
	// 		result = "it is NOT palindrome"
	// 	}
	// }
	// fmt.Println(data)
	// fmt.Println(result)

	show := palindrom("MALAM")
	fmt.Println(show)
}

func palindromV2(data string) string {
	var rev string
	result = "its not a palindrome
	for i := 0; i < len(data); i++ {
		rev += string(data[len(data) - 1 - i])
	}
	if rev == data {
		result = "its palindrome"
	}
	return result
}
func palindrom(data string) string {
	var result string
	for i := 0; i < len(data)/2; i++ {
		dataReserve := len(data) - 1 - i

		if data[i] == data[dataReserve] {
			result = "it is palindrome"
		}else {
			result = "it is NOT palindrome"
		}
	}
	fmt.Println(data)
	return result
}
