package main

import "fmt"

func main() {
	// stampo la lista dei numeri primi da 0 a n
START:
	fmt.Println("inserisci il numero")
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("il numero deve essere maggiore di 1")
		goto START
	}
	var isPrime bool = true
	for i := n; i >= 2; i-- {
		for j := n - 1; j >= 2; j-- {
			if n%j == 0 {
				isPrime = false
			}

		}
		fmt.Println(i)
		if isPrime {
			x := []int{}
			x = append(x, i)
		}

	}
	fmt.Println(isPrime)
}
