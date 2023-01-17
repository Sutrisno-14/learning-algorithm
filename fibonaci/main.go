package main

import (
	"fmt"

	// "github.com/go-playground/locales/ln"
)

func fibonaci(p int) {
	var dataFibonacci []int
	n :=1
	n1 :=1
	n2 := 0

	for i :=0; i<p; i++ {
		// push hasil ke array
		dataFibonacci = append(dataFibonacci, n)
		n = n1+n2 
		fmt.Println("n :",n)
		n2 = n1 // n1 bakal menjadi n2 atau angka sebelumny
		fmt.Println("n2 :",n2)
		n1 = n // n bakal menjadi n1 atau angka sekarang
		fmt.Println("n1 :",n1)

	}
	fmt.Println(dataFibonacci)
}
func main() {
	var p int
	fmt.Printf("input angka : ")
	fmt.Scan(&p)
	fibonaci(p)
}