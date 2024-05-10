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

	stock.CalculateTotalShareAmount(o.Quantity)
	stock.CalculateWeightedAvgPrice(operationIndex)
	grossProfit := o.calculateGrossProfit(stock.weightedAvgPrice)

	isTotalAmountTaxable := o.TotalAmount > MinTaxableAmount
	hasLosses := stock.losses < 0
	isDeductable := isTotalAmountTaxable && hasLosses
	o.GetNetProfit(isDeductable, grossProfit, stock.losses)

	if o.NetProfit < 0 {
		stock.DeductLoss(grossProfit)
	}

	if isDeductable {
		o.NetProfit = netProfitCalculation(grossProfit, stock.losses)
		stock.DeductLoss(grossProfit)
	} else {
		o.NetProfit = grossProfit
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

func (o *Operation) GetNetProfit(isDeductable bool, grossProfit, losses float32) {
	if isDeductable {
		o.NetProfit = netProfitCalculation(grossProfit, losses)
	} else {
		o.NetProfit = grossProfit
	}
}

func (o *Operation) calculateGrossProfit(weightedAvgPrice float32) float32 {
	return (float32(o.Quantity) * o.UnitCost) - (float32(o.Quantity) * weightedAvgPrice)
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
	stoks := strings.Split(data, "]")

	var items [][]Operation

	for i := 0; i < len(stoks)-1; i++ {
		jsonStock := stoks[i] + "]"
		var partialOperations []Operation
		var operations []Operation

		// Convert the string with the list of operation into a struct
		if err := json.Unmarshal([]byte(strings.TrimSpace(jsonStock)), &partialOperations); err != nil {
			log.Fatalf("OPERATION L40: %s", err)
		}

		// Call the NewOperation method to create the operations for each stock
		for _, op := range partialOperations {
			operation := NewOperation(op.Type, op.UnitCost, op.Quantity)
			operations = append(operations, operation)
		}

		items = append(items, operations)
	}

	return items
}
