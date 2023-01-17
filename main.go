package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var prima []int
	for i := 1; i <= 10; i++ {
		if isPrime(i) {
			prima = append(prima, i)
		}
	}
	fmt.Println(prima)
}
