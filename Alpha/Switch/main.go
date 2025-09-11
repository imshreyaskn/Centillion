package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func math() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the method:")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(method)

	var val1 int
	var val2 int

	fmt.Print("Enter the first val:")
	_, _ = fmt.Scanln(&val1)
	fmt.Print("Enter the second val:")
	_, _ = fmt.Scanln(&val2)

	switch strings.ToLower(method) { // make it case-insensitive
	case "add":
		fmt.Println("Result:", val1+val2)
	case "subtract":
		fmt.Println("Result:", val1-val2)
	case "multiply":
		fmt.Println("Result:", val1*val2)
	case "divide":
		if val2 == 0 {
			fmt.Println("Cannot divide by zero!")
		} else {
			fmt.Println("Result:", val1/val2)
		}
	default:
		fmt.Println("Unknown method:", method)
	}
}

func main() {

	math()

}
