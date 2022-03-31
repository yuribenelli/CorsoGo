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

	prs1 := person{
		first_name:    "James",
		last_name:     "Bond",
		fav_ice_cream: []string{"chocolate", "banana", "hazelnut"},
	}
	prs2 := person{
		first_name:    "Miss",
		last_name:     "Moneypenny",
		fav_ice_cream: []string{"strawberry", "chocolate", "Cream"},
	}
	fmt.Println(prs1.first_name, prs1.last_name)
	for i, v := range prs1.fav_ice_cream {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Println(prs2.first_name, prs2.last_name)
	for i, v := range prs2.fav_ice_cream {
		fmt.Printf("%v\t%v\n", i, v)
	}
	//EX2
	fmt.Println("ESERCIZIO 2")
	myMap := map[string]person{
		prs1.last_name: prs1,
		prs2.last_name: prs2,
	}
	fmt.Println("")
	fmt.Println(myMap)
	fmt.Println("")

	for _, v := range myMap {
		fmt.Println(v.first_name, v.last_name)

		for i, val := range v.fav_ice_cream {
			fmt.Println(i, val)

		}
	}
	//EX3
	fmt.Println("ESERCIZIO 3")
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}
	myTruck := truck{
		vehicle: vehicle{
			doors: 5,
			color: "blue",
		},
		fourWheel: true,
	}
	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println("")
	fmt.Println(myTruck)
	fmt.Println(mySedan)
	fmt.Println("")
	fmt.Printf("%T\n", mySedan)
	fmt.Printf("Color: \t\t%v\n", mySedan.color)
	fmt.Printf("Doors: \t\t%v\n", mySedan.doors)
	fmt.Printf("Luxury: \t%v\n", mySedan.luxury)
	fmt.Println("")
	fmt.Printf("%T\n", myTruck)
	fmt.Printf("Color: \t\t%v\n", myTruck.color)
	fmt.Printf("Doors: \t\t%v\n", myTruck.doors)
	fmt.Printf("Four wheel: \t%v\n", myTruck.fourWheel)

	//EX4
	fmt.Println("ESERCIZIO 4")
	fmt.Println("")

	miaStruttura := struct {
		Nome    string
		Cognome string
		Anni    int
	}{
		Nome:    "Yuri",
		Cognome: "Benelli",
		Anni:    38,
	}

	fmt.Println(miaStruttura)
}
