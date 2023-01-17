package main

import "fmt"

func Server(arr []int32) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 || arr[i]%2 == -1 {
			fmt.Println(arr[i])
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Println(arr[i])
		}
	}
}

func main() {
	var n, r int
	var arr []int32

	fmt.Printf("input panjang array dan isi array : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&r)
		arr = append(arr, int32(r))
	}

	Server(arr)
}