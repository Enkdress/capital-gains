package main

import (
	"reflect"
	"testing"
)

func TestCalculateGrossProfit(t *testing.T) {
	var avgPrice float32 = 10.0
	var unitCost float32 = 20.0
	var quantity int32 = 100
	operation := NewOperation(SellOperation, unitCost, quantity)

	grossProfit := operation.calculateGrossProfit(avgPrice)
	expected := 1000

	if grossProfit != float32(expected) {
		t.Errorf("\nFailed gross profit calculation:\nExpected: %v\nRecieved: %v\n", expected, grossProfit)
	}
}

func TestCalculateNetProfit(t *testing.T) {
	var avgPrice float32 = 10.0
	var losses float32 = 0.0
	var unitCost float32 = 20.0
	var quantity int32 = 100
	operation := NewOperation(SellOperation, unitCost, quantity)

	grossProfit := operation.calculateGrossProfit(avgPrice)
	netProfit := operation.calculateNetProfit(grossProfit, losses)
	expected := 1000

	if netProfit != float32(expected) {
		t.Errorf("\nFailed net profit calculation:\nExpected: %v\nRecieved: %v\n", expected, netProfit)
	}

	avgPrice = 10.0
	losses = 0.0
	unitCost = 10.0
	quantity = 100
	operation = NewOperation(SellOperation, unitCost, quantity)

	grossProfit = operation.calculateGrossProfit(avgPrice)
	netProfit = operation.calculateNetProfit(grossProfit, losses)
	expected = 0

	if netProfit != float32(expected) {
		t.Errorf("\nFailed net profit calculation:\nExpected: %v\nRecieved: %v\n", expected, netProfit)
	}

}

func TestCalculateTaxOverProfit(t *testing.T) {
}

func TestCase1(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case1Input.GetTaxes(), Case1Result)
	if !isTaxEqual {
		t.Errorf("\nCase 1 failed:\nExpected: %v\nRecieved: %v", Case1Result, Case1Input.GetTaxes())
	}
}

func TestCase2(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case2Input.GetTaxes(), Case2Result)
	if !isTaxEqual {
		t.Errorf("\nCase 2 failed:\nExpected: %v\nRecieved: %v", Case2Result, Case2Input.GetTaxes())
	}
}

func TestCase3(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case3Input.GetTaxes(), Case3Result)
	if !isTaxEqual {
		t.Errorf("\nCase 3 failed:\nExpected: %v\nRecieved: %v", Case3Result, Case3Input.GetTaxes())
	}
}

func TestCase4(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case4Input.GetTaxes(), Case4Result)
	if !isTaxEqual {
		t.Errorf("\nCase 4 failed:\nExpected: %v\nRecieved: %v", Case4Result, Case4Input.GetTaxes())
	}
}

func TestCase5(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case5Input.GetTaxes(), Case5Result)
	if !isTaxEqual {
		t.Errorf("\nCase 5 failed:\nExpected: %v\nRecieved: %v", Case5Result, Case5Input.GetTaxes())
	}
}

func TestCase6(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case6Input.GetTaxes(), Case6Result)
	if !isTaxEqual {
		t.Errorf("\nCase 6 failed:\nExpected: %v\nRecieved: %v", Case6Result, Case6Input.GetTaxes())
	}
}

func TestCase7(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case7Input.GetTaxes(), Case7Result)
	if !isTaxEqual {
		t.Errorf("\nCase 7 failed:\nExpected: %v\nRecieved: %v", Case7Result, Case7Input.GetTaxes())
	}
}

func TestCase8(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case8Input.GetTaxes(), Case8Result)
	if !isTaxEqual {
		t.Errorf("\nCase 8 failed:\nExpected: %v\nRecieved: %v", Case8Result, Case8Input.GetTaxes())
	}
}

func TestCase9(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case9Input.GetTaxes(), Case9Result)
	if !isTaxEqual {
		t.Errorf("\nCase 9 failed:\nExpected: %v\nRecieved: %v", Case9Result, Case9Input.GetTaxes())
	}
}

var Case1Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 100),
	NewOperation(SellOperation, 15.00, 50),
	NewOperation(SellOperation, 15.00, 50),
}
var Case1Input = NewStock(Case1Operations)

var Case1Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
}

var Case2Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(SellOperation, 20.00, 5000),
	NewOperation(SellOperation, 5.00, 5000),
}
var Case2Input = NewStock(Case2Operations)
var Case2Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 10000.00},
	{"tax": 0.00},
}

var Case3Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(SellOperation, 5.00, 5000),
	NewOperation(SellOperation, 20.00, 3000),
}
var Case3Input = NewStock(Case3Operations)
var Case3Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 1000.00},
}

var Case4Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(BuyOperation, 25.00, 5000),
	NewOperation(SellOperation, 15.00, 10000),
}
var Case4Input = NewStock(Case4Operations)
var Case4Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
}

var Case5Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(BuyOperation, 25.00, 5000),
	NewOperation(SellOperation, 15.00, 10000),
	NewOperation(SellOperation, 25.00, 5000),
}
var Case5Input = NewStock(Case5Operations)
var Case5Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 10000.00},
}

var Case6Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(SellOperation, 2.00, 5000),
	NewOperation(SellOperation, 20.00, 2000),
	NewOperation(SellOperation, 20.00, 2000),
	NewOperation(SellOperation, 25.00, 1000),
}
var Case6Input = NewStock(Case6Operations)
var Case6Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 3000.00},
}

var Case7Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(SellOperation, 2.00, 5000),
	NewOperation(SellOperation, 20.00, 2000),
	NewOperation(SellOperation, 20.00, 2000),
	NewOperation(SellOperation, 25.00, 1000),
	NewOperation(BuyOperation, 20.00, 10000),
	NewOperation(SellOperation, 15.00, 5000),
	NewOperation(SellOperation, 30.00, 4350),
	NewOperation(SellOperation, 30.00, 650),
}
var Case7Input = NewStock(Case7Operations)
var Case7Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 3000.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 3700.00},
	{"tax": 0.00},
}

var Case8Operations = []Operation{
	NewOperation(BuyOperation, 10.00, 10000),
	NewOperation(SellOperation, 50.00, 10000),
	NewOperation(BuyOperation, 20.00, 10000),
	NewOperation(SellOperation, 50.00, 10000),
}
var Case8Input = NewStock(Case8Operations)
var Case8Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 80000.00},
	{"tax": 0.00},
	{"tax": 60000.00},
}

var Case9Operations = []Operation{
	NewOperation(BuyOperation, 5000.00, 10),
	NewOperation(SellOperation, 4000.00, 5),
	NewOperation(BuyOperation, 15000.00, 5),
	NewOperation(BuyOperation, 4000.00, 2),
	NewOperation(BuyOperation, 23000.00, 2),
	NewOperation(SellOperation, 20000.00, 1),
	NewOperation(SellOperation, 12000.00, 10),
	NewOperation(SellOperation, 15000.00, 3),
}
var Case9Input = NewStock(Case9Operations)
var Case9Result = []map[string]float32{
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 0.00},
	{"tax": 1000.00},
	{"tax": 2400.00},
}
