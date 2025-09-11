package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bill struct {
	id    int
	items map[string]float64
	tip   float64
	total float64
}

func newBill(id int) Bill {
	b := Bill{
		id:    id,
		items: map[string]float64{},
		tip:   0,
		total: 0,
	}

	return b
}

func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}
func (b *Bill) updateTotal() {
	var total float64 = 0
	for _, v := range b.items {
		total += v
	}

	b.total = total + b.tip
}

func (b *Bill) addItem() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter the item name:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Enter the item pice:")
		pricestr, _ := reader.ReadString('\n')
		pricestr = strings.TrimSpace(pricestr)
		price, err := strconv.ParseFloat(pricestr, 64)

		if err != nil {
			fmt.Println("Invalid Price Try Again")
			continue
		}

		b.items[name] = price

		fmt.Println("Need to add more item (yes/no):")

		cont, _ := reader.ReadString('\n')
		cont = strings.TrimSpace(cont)

		if strings.ToLower(cont) == "no" {
			break
		}

	}
}

func (b *Bill) format() string {
	fs := fmt.Sprintf("Bill ID: %d\n", b.id)
	fs += "Bill Breakdown:\n"
	for k, v := range b.items {
		fs += fmt.Sprintf("%-20v ...$%v\n", k+":", v)
	}
	fs += fmt.Sprintf("%-20v ...$%.2f\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-20v ...$%.2f", "total:", b.total)
	return fs
}

func (b *Bill) save() {
	os.MkdirAll("bills", 0775)

	filename := fmt.Sprintf("bills/bill_%v.txt", b.id)

	data := []byte(b.format())

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved successfully to:", filename)

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Bill id: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	bill := newBill(id)
	bill.addItem()

	fmt.Print("Enter tip amount: ")
	var tip float64
	fmt.Scanln(&tip)
	bill.updateTip(tip)
	bill.updateTotal()

	bill.save()

}
