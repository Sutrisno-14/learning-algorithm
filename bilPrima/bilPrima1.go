package main

import "fmt"

// func isPrime(n int) bool {
// 	for i := 2; i < n;i++ {
// 		if n%2 ==0 {
// 			return false
// 		}
// 	}
// 	return true
// }
// func showPrime(max int) {
// 	for i := 2; i <= max; i++ {
// 		if isPrime(i) {
// 			fmt.Println(i)
// 		}
// 	}
// }

func isPrimee(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func showPrime(max int) {
	for i := 2; i <= max; i++ {
		if isPrimee(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	showPrime(10)
}