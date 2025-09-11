package main

import (
	"fmt"
)

func main() {
	Mary_Bill := newBill(1024)
	//Mary_Bill.addItems("Soup", 10)
	//Mary_Bill.addItems("Noodles", 15)
	//Mary_Bill.updateTip(10)
	fmt.Println(Mary_Bill.format())
}
