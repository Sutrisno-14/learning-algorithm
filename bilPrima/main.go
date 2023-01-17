package main

import "fmt"

func Prima(n int) {
	
	var bilPrima []int

	for i := 1; i <= n; i++ {
		var bill int = 0 //state pengecekan apabila nilai habis di bagi 1 dan nilai itu sendiri
		for j := 1; j<=i; j++ {
			if i%j == 0 {
				bill = bill + 1
			}
		}
		fmt.Println(bill)
		if bill == 2 {
			bilPrima = append(bilPrima, i)
		}
	}
	fmt.Println(bilPrima)
}

func main() {
	var n int
	fmt.Printf("input banyak data : ")
	fmt.Scan(&n)
	Prima(n)
}