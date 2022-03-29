package main

import "fmt"

func main() {
	//Using two methods for iteration

	count := 100
	for i := 0; i < count; i++ {
		fmt.Printf("%v\n", i)
	}

	//single condition
	i := 0
	for i < 100 {
		fmt.Printf("%v\n", i)
		i++

	}

}
