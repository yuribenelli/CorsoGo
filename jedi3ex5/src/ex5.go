package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 0
	countDiv3 := 0
	maxLoop := 100
	for i < maxLoop {
		rnd := rand.Intn(100)
		fmt.Printf("%v\t", rnd)
		if rnd%3 == 0 {
			countDiv3++
		}
		i++

	}
	fmt.Printf("\n di %v numeri generati, %v sono divisibili per 3", maxLoop, countDiv3)
}
