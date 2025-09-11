package main

import "fmt"

type Bill struct {
	id    int
	items map[string]float64
	tip   float64
}

func newBill(id int) Bill {
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

	fs += fmt.Sprintf("%-25v ...%v\n", "tip:", b.tip)

	//total

	fs += fmt.Sprintf("%-25v ...%v", "total:", total+b.tip)

	return fs
}

func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}
