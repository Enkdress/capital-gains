package main

// Fixed tax at 20% of the profit
const TaxRate = 0.2
const MinTaxableAmount = 20_000

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
