package main

import "testing"

func TestCalculateAvgPrice(t *testing.T) {
	var stock Stock
	stock.weightedAvgPrice = 10.0 // asumming the first buy unit cost is 10

	avgPrice := stock.CalculateAvgPrice(100, 50, 10.0)
	if avgPrice != 10.0 {
		t.Errorf("Expected avg price 10.0, got %f", avgPrice)
	}

	stock.operations = append(stock.operations, Operation{UnitCost: 15.0})
	avgPrice = stock.CalculateAvgPrice(100, 20, 20.0)
	var expectedPrice float32 = (100*10.0 + 20*20.0) / 120.0
	if avgPrice != expectedPrice {
		t.Errorf("Expected avg price %f, got %f", expectedPrice, avgPrice)
	}
}
