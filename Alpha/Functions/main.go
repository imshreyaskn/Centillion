package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 10
	y := 20
	names := []string{"Santhosh", "Shreyas", "Niranjan"}
	fmt.Println(add(x, y))
	fmt.Println(sub(x, y))
	greet(names, greetHelper)
	firstInitial, secondIntial := getInitial("Adithya Varma")
	fmt.Printf("First Initial: %v. Second Initial: %v.", firstInitial, secondIntial)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//Function As Arguments

func greet(names []string, helperFunc func(name string)) {
	for _, name := range names {
		helperFunc(name)
	}
}
func greetHelper(name string) {
	fmt.Printf("Welcome to da club, %v.\n", name)
}

//Multiple Returns

func getInitial(n string) (string, string) {
	s := strings.ToUpper(n)
	splittedName := strings.Split(s, " ")

	var initials []string

	for _, values := range splittedName {
		initials = append(initials, values[0:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}
}
