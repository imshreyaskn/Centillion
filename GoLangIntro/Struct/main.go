package main

import (
	"fmt"
)

func main() {
	Mary_Bill := newBill(1024)
	Mary_Bill.updateTip(10)
	fmt.Println(Mary_Bill.format())
}
