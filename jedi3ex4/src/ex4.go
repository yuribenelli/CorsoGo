package main

import "fmt"

func main() {

	birthDate := 1984
	currentYear := 2022

	for {

		if currentYear < birthDate {
			break
		}
		fmt.Printf("%v\n", birthDate)
		birthDate++

	}
}
