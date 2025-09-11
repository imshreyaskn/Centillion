package main

import (
	"fmt"
	"strings"
)

func main() {

	//11-09-2025
	//Declaration & Initialization

	age := 35
	name := "Sean"
	points := 25.35

	//Output Formating

	fmt.Printf("My age is %v. \n", age)
	fmt.Printf("My age type is %T. \n", age)
	fmt.Printf("My age in queates: %q. \n", name)
	fmt.Printf("Points: %.2f \n", points)

	//Loops Control Statements

	//Arrays

	var ages [3]int = [3]int{20, 25, 30}
	names := [3]string{"Mario", "Yoshi", "Lucy"}

	fmt.Println(len(ages))

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//Slices

	var values []int = []int{10, 23, 55}

	values = append(values, 85)

	valueRange := values[1:2]

	valueRange[2] = 10

	//Standard Library - Strings

	fmt.Println(strings.Contains("Heyy", "ey"))
	fmt.Println(strings.ReplaceAll("Heyy", "y", "i"))
	fmt.Println(strings.ToUpper("Caps"))
	fmt.Println(strings.Index("Caps", "aps"))

}
