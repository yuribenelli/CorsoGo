package main

import "fmt"

func main() {
	count := 150000
	for i := 1; i < count; i++ {
		a := i * i
		fmt.Printf("%v\t%v\n", i, a)
		fmt.Printf("%T\n", i)

	}

}
