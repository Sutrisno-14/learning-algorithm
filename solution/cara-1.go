package main

import (
	"fmt"
	"sort"
)


func solution(A []int) int {

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

func arr() []int{
	var n, arr int
	var a []int

	fmt.Println("Input panjang array dan angka dalam array : ")
	fmt.Scan(&n)
	
	//tambahkan angka di variabel arr kedalam array A
	for i := 0; i < n; i++ {
		fmt.Scan(&arr)
		a = append(a, arr)
	}
	return a
}

func main() {
	// var n, arr int
	// var A []int
	// fmt.Printf("Input panjang array dan isi array:")
	// fmt.Scan(&n)
	// fmt.Println("Input isi dari array (sesuai panjang array) :")
	// fmt.Scan(&arr)
	// A := []int{1,3,6,4,1,2}
	// fmt.Println(solution(A))
	// A := []int{1, 2, 3}
	// fmt.Println(solution(A)) // Output: 4
	A := arr()
	fmt.Println(solution(A))

	// A := []int{-1, -3}
	// fmt.Println(solution(A)) // Output: 1

	
}