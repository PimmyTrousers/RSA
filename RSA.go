package main

import "fmt"

func modInverse(e int, pn int) int {
	e = e % pn
	for d := 1; d < pn; d++ {
		if (e*d)%pn == 1 {
			return d
		}
	}
	return 0
}

func main() {
	fmt.Println("Enter 2 prime numbers and one number that is coprime: \n ")

}
