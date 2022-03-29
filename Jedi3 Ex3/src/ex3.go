package main

import "fmt"

func main() {
	birthDate := 1984
	currentYear := 2022

	for birthDate < currentYear {
		fmt.Printf("%v\n", birthDate)
		birthDate++

	}
}
