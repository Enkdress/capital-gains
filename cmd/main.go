package main

import (
	"fmt"
)

func main() {
	text := GetStinData()
	stockOperations := ParseOperations(text)
	stocks := InitializeStocks(stockOperations)
	for _, stock := range stocks {
		fmt.Println(stock.taxes)

	}
}
