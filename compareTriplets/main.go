package main

import "fmt"

func main() {
	a := []int{17,28,30}
	b := []int{99,16,8}
	var alice, boby int
	var data []int

	for i :=0; i < len(a); i++ {
		if a[i] > b[i] {
			alice+=1
		}else if a[i] < b[i] {
			boby+=1
		}
	}
	data = append(data, alice,boby)
	fmt.Println(data)
}