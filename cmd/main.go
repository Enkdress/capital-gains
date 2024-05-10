package main

import (
	"fmt"
)

// For each set of operations create a stock struct
func InitializeStocks(stockOperations [][]Operation) []Stock {
	var stocks []Stock
	for i := 0; i < len(stockOperations); i++ {
		stock := NewStock(stockOperations[i])
		stocks = append(stocks, stock)
	}

	return stocks
}

func main() {
	text := GetStinData()
	stockOperations := ParseOperations(text)
	stocks := InitializeStocks(stockOperations)
	for _, stock := range stocks {
		fmt.Println(stock.GetTaxes())
	}
}
