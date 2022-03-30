package main

import (
	"fmt"
)

func main() {
	// EX1 EX2   *******
	fmt.Println("ESERCIZIO 1 E 2:")
	ds := [5]int{}
	ds[0] = 12
	ds[1] = 23
	ds[2] = 6
	ds[3] = 1887
	ds[4] = 14
	//Printing array values
	fmt.Println(ds)
	//Printing Array Type
	fmt.Printf("%T\n\n", ds)
	//Range cycle over ds. Print values and indexes
	for i, v := range ds {
		fmt.Println(i, v)
	}
	//Creating a slice of TYPE int
	ss := []int{10, 34, 54, 198, 97}
	fmt.Println("")
	//Range cycle over ss. Print values
	for i, v := range ss {
		fmt.Println(i, v)
	}
	fmt.Printf("\n%T\n", ss)

	// EX3 *******
	fmt.Println("ESERCIZIO 3:")
	data := []int{}
	//populating slice
	for i := 42; i < 52; i++ {
		data = append(data, i)
	}
	fmt.Printf("%v\n", data)
	//Printing slices of the data slice
	fmt.Println(data[:5], data[5:], data[2:7], data[1:6])

	// EX4 *******
	fmt.Println("ESERCIZIO 4:")
	data1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(data1)
	data1 = append(data1, 52)
	fmt.Println(data1)
	data1 = append(data1, 53, 54, 55)
	fmt.Println(data1)
	y := []int{56, 57, 58, 59, 60}
	data1 = append(data1, y...)
	fmt.Println(data1)

	//EX5 *****
	fmt.Println("ESERCIZIO 5:")
	data5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	report := []int{}
	report = append(data5[:3], data5[6:]...)
	fmt.Println(report)

	//EX6
	fmt.Println("ESERCIZIO 6:")
	states := make([]string, 0, 50)
	fmt.Printf("%v\n%v\n", len(states), cap(states))
	states = append(states, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`,
		`Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`,
		`Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`,
		`Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`,
		`New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`,
		`Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`,
		`Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)
	fmt.Println(states)
	fmt.Printf("%v\n%v\n", len(states), cap(states))
	for i, v := range states {
		fmt.Printf("%v\t%v\n", i, v)
	}

	//EX7
	fmt.Println("ESERCIZIO 7:")
	sl1 := []string{"James", "Bond", "Shaken, not stirred"}
	sl2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	sum := [][]string{sl1, sl2}
	fmt.Println(sum)
	for i, v := range sum {
		fmt.Println(i)
		for j, val := range v {
			fmt.Printf("index: %v\t value: %v\n", j, val)
		}

	}

	//EX8
	fmt.Println("ESERCIZIO 8:")
	pref := map[string][]string{
		"bond_james":      {"shaken, not stirred", "martinis", "women"},
		"Moneypenny_miss": {"James Bond", "literature", "Computer science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(pref)
	for j, v := range pref {
		fmt.Println("record:", j)
		for z, t := range v {
			fmt.Println("\t", z, t)

		}
	}
	//EX9
	fmt.Println("ESERCIZIO 9:")
	pref["Q_Mister"] = []string{"Technologies", "cars", "guns"}
	for j, v := range pref {
		fmt.Println("record:", j)
		for z, t := range v {
			fmt.Println("\t", z, t)
		}
	}

	//EX10
	fmt.Println("ESERCIZIO 8:")
	delete(pref, "Q_Mister")
	for j, v := range pref {
		fmt.Println("record:", j)
		for z, t := range v {
			fmt.Println("\t", z, t)
		}
	}
}
