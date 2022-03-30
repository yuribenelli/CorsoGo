package main

import (
	"fmt"
)

func main() {
	//EX1
	fmt.Println("ESERCIZIO 1")
	type person struct {
		first_name    string
		last_name     string
		fav_ice_cream []string
	}

	p1 := person{
		first_name:    "James",
		last_name:     "Bond",
		fav_ice_cream: []string{"chocolate", "banana", "hazelnut"},
	}
	p2 := person{
		first_name:    "Miss",
		last_name:     "Moneypenny",
		fav_ice_cream: []string{"strawberry", "chocolate", "Cream"},
	}
	fmt.Println(p1.first_name, p1.last_name)
	for i, v := range p1.fav_ice_cream {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Println(p2.first_name, p2.last_name)
	for i, v := range p2.fav_ice_cream {
		fmt.Printf("%v\t%v\n", i, v)
	}
}
