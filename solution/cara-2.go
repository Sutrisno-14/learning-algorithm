package main

import (
	"fmt"
	"sort"
)

func solutions(A []int) int {

	// array A diurutkan secara ascending
	sort.Ints(A)
	fmt.Println("sort", A)
	//inisialisasi variabel hasil dengan nilai 1

	result := 1

	for i := 0; i < len(A); i++ {
		// jika A[i] > hasil, maka keluar dari perulangan
		// fmt.Println("hasil ===", result)
		if A[i] > result {
			break
		}

		//jika A[i] == hasil, tambahkan hasil dengan 1
		// fmt.Println("ini result ===",result)
		fmt.Println("A[",i,"]", A[i] ,"==", result)
		if A[i] == result {
			result = result + 1
			fmt.Println("result =", result)
		}
		// fmt.Println("result",result)
	}
	return result
}

func main() {
	var n, arr int
	var A []int

	fmt.Printf("input panjang array dan isi array : ")
	fmt.Scan(&n)
	// append value arr ke dalam variabel array A
	for i := 0; i < n; i++ {
		fmt.Scan(&arr)
		A = append(A, arr)
	}

	//Menjalankan function solution(parameter)
	fmt.Println(solutions(A))
}