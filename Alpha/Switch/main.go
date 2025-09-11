package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func math() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the method:")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(method)

	val1str, _ := reader.ReadString('\n')
	val1str = strings.TrimSpace(val1str)
	val1, _ := strconv.ParseFloat(val1str, 64)

	val2str, _ := reader.ReadString('\n')
	val2str = strings.TrimSpace(val2str)
	val2, _ := strconv.ParseFloat(val2str, 64)

	switch strings.ToLower(method) {
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
