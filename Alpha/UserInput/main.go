package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Bill Id:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	b := newBill(id)

	return b

}

func main() {
	Mary_Bill := createBill()
	Mary_Bill.updateTip()
	fmt.Println(Mary_Bill.format())
}
