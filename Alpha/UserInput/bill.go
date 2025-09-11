package main

import (
	"fmt"
)

type Bill struct {
	id    string
	items map[string]float64
	tip   float64
}

func newBill(id string) Bill {
	b := Bill{
		id:    id,
		items: map[string]float64{},
		tip:   0.0,
	}
	return b
}

func (b *Bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	//list items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	//Tip

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	//total

	fs += fmt.Sprintf("%-25v ...$%v\n", "total:", total+b.tip)

	return fs
}

func (b *Bill) updateTip() {

	var tip float64
	fmt.Print("Enter an tip: ")
	_, err := fmt.Scanln(&tip)
	if err != nil {
		fmt.Println("err")
		return
	}
	b.tip = tip
}

func (b *Bill) addItems(name string, price float64) {
	b.items[name] = price
}
