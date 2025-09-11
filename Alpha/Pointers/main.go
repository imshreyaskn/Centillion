package main

import (
	"fmt"
)

func main() {
	name := "Tifa"
	mem := &name
	fmt.Printf("The memeory adress of the variable name is %v.", mem)
	fmt.Printf("The value of the variable name is %v.", *mem)

	//pass by reference

	updateName(mem)

	fmt.Printf("Value after pass by reference : %v.\n", name)

}

func updateName(s *string) {
	*s = "Lucy"
}
