package main

import (
	"encoding/json"
	"log"
	"strings"
)

const (
	BuyOperation  string = "buy"
	SellOperation string = "sell"
)

type Operation struct {
	Type        string  `json:"operation"`
	UnitCost    float32 `json:"unit-cost"`
	Quantity    int32   `json:"quantity"`
	TotalAmount float32
	NetProfit   float32
}

func (o *Operation) CalculateTaxes(stock *Stock, operationIndex int) map[string]float32 {
	// Only sell operations has taxes
	if o.Type == BuyOperation {
		stock.totalQuantity = stock.totalQuantity + o.Quantity
		return map[string]float32{"tax": 0.00}
	}

	// This is the amount of shares of the stock
	stock.totalQuantity = stock.totalQuantity - o.Quantity
	stock.CalculateWeightedAvgPrice(operationIndex)

	grossProfit := o.calculateGrossProfit(stock.weightedAvgPrice)
	isTotalAmountTaxable := o.TotalAmount > MinTaxableAmount
	hasLosses := stock.losses < 0
	isDeductable := isTotalAmountTaxable && hasLosses

	if isDeductable {
		o.NetProfit = netProfitCalculation(grossProfit, stock.losses)
		stock.losses = stock.losses + grossProfit // Deduct the profit to the losses (like paying my debt)
	}

	isTaxable := isTotalAmountTaxable && o.NetProfit > 0

	// Losses do not pay taxes
	if grossProfit < 0 {
		stock.AcummulateLosses(grossProfit)
		return map[string]float32{"tax": 0.00}
	}

	if isTaxable {
		taxValue := taxCalculation(o.NetProfit)

		if isDeductable {
			stock.AcummulateProfit(grossProfit + stock.losses)
		}

		stock.ResetProfit()

		return map[string]float32{"tax": taxValue}
	}

	return map[string]float32{"tax": 0.00}
}

func netProfitCalculation(grossProfit, profit float32) float32 {
	value := grossProfit + profit
	netProfit := grossProfit

	if value < 1 {
		netProfit = 0
	} else {
		netProfit = value
	}

	return netProfit
}

func taxCalculation(netProfit float32) float32 {
	return netProfit * TaxRate

}

func (o *Operation) calculateGrossProfit(weightedAvgPrice float32) float32 {
	return (float32(o.Quantity) * o.UnitCost) - (float32(o.Quantity) * weightedAvgPrice)
}

func NewOperation(operationType string, unitCost float32, quantity int32) Operation {
	operation := Operation{
		Type:        operationType,
		UnitCost:    unitCost,
		Quantity:    quantity,
		TotalAmount: unitCost * float32(quantity),
		NetProfit:   0,
	}

	return operation
}

func ParseOperations(data string) [][]Operation {
	parts := strings.Split(data, "]")

	var items [][]Operation

	for i := 0; i < len(parts)-1; i++ {
		jsonString := parts[i] + "]"
		var operations []Operation

		if err := json.Unmarshal([]byte(strings.TrimSpace(jsonString)), &operations); err != nil {
			log.Fatalf("OPERATION L40: %s", err)
		}
		items = append(items, operations)
	}

	return items
}
