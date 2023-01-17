package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var a string
	fmt.Printf("Masukan kata : ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\nScanner :\n", scanner)
	scanner.Scan()
	a := scanner.Text()
	// fmt.Scanln(&a)
	coba := test(a)
	fmt.Println(coba)
	fmt.Println(a)

}

func test(data string) string {
	for i := 0; i < len(data)/2; i++ {
		if data[i] == data[len(data)-1-i] {
			return "palindrome"
		}
	}
	return "bukan palindrome"
}
