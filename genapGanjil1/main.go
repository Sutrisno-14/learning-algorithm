package main

import "fmt"

func genapGanjil(n int) {
	var ganjil, genap []int
	for i := 1; i <= n; i ++ {
		if i%2 == 0 {
			genap = append(genap, i)
		}else {
			ganjil = append(ganjil, i)
		}
	}
	fmt.Println("Genap : ", genap)
	fmt.Println("Ganjil : ", ganjil)
}

func main() {
	var n int
	fmt.Printf("input data ke-n : ")
	fmt.Scan(&n)
	genapGanjil(n)
	
}
