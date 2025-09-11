package main

import (
	"fmt"
)

func main() {

	//While Loop

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	//For Loop
	for y := 0; y < 5; y++ {
		fmt.Println(y)
	}

	//For-Range Loop

	names := []string{"santhosh", "shreyas", "niranjan"}
	for index, value := range names {
		fmt.Printf("The Value: %v. The Index: %v \n", value, index)
	}

	// Use underscore '_' to discard unwanted variable

	//Slice Modification using For-Range loop

	for indx, _ := range names {
		names[indx] = "None"
	}

	//Infinte loop

	z := 0
	for {
		if z > 5 {
			break
		}
		fmt.Println(z)
		z++
	}

	//Contine and Break key

	m := 0
	for {
		if m == 3 {
			continue //skips the num 3
		} else if m == 5 {
			break //Breaks the loop when m = 5
		}
		fmt.Println(m)
		m++
	}

}
