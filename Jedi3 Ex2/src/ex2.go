package main

import "fmt"

func main() {

	for i := 1; i <= 25; i++ {
		fmt.Println(i)
		for j := 1; j < 4; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
