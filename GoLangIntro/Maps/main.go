package main

import "fmt"

func main() {
	items := map[string]float64{
		"Soap":    20.00,
		"Powder":  50.00,
		"Paste":   10.00,
		"Shampoo": 60.00,
	}

	fmt.Printf("Price of Soap: %v.\n", items["Soap"])
	fmt.Println("Price of All the items")

	for key, value := range items {
		fmt.Printf("Price of %.v is %v.\n", key, value)
	}

	//Defaulting to base price

	for key, _ := range items {
		items[key] = 0.00
	}

}
