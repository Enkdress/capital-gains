package main

type Stock struct {
	weightedAvgPrice float32
	operations       []Operation
	totalQuantity    int32
	profit           float32
	losses           float32
}

func (s *Stock) UpdateTotalShareCount(quantity int32, operationType string) {
	if operationType == BuyOperation {
		s.totalQuantity = s.totalQuantity + quantity
	} else {
		s.totalQuantity = s.totalQuantity - quantity
	}
}

func (s *Stock) AcummulateProfit(grossProfit float32) {
	s.profit = s.profit + grossProfit
}

func (s *Stock) AcummulateLosses(grossProfit float32) {
	s.losses = s.losses + grossProfit
}

func (s *Stock) ResetProfit() {
	if s.totalQuantity == 0 {
		s.profit = 0
	}
}

func (s *Stock) CalculateWeightedAvgPrice(currentOperation int) {
	currentQuantity := s.operations[0].Quantity
	arrInit := 0

	for i, operation := range s.operations[arrInit:currentOperation] {
		if i == 0 {
			s.weightedAvgPrice = operation.UnitCost
			continue
		}

		if operation.Type == SellOperation {
			currentQuantity = currentQuantity - operation.Quantity
			continue
		}

		currentWeightedAvgPrice := ((float32(currentQuantity) * s.weightedAvgPrice) +
			(float32(operation.Quantity) * operation.UnitCost)) /
			float32(currentQuantity+operation.Quantity)

		s.weightedAvgPrice = currentWeightedAvgPrice
		currentQuantity = currentQuantity + operation.Quantity
	}
}

func (s *Stock) GetTaxes() []map[string]float32 {
	taxes := make([]map[string]float32, 0, 0)
	for i, operation := range s.operations {
		s.UpdateTotalShareCount(operation.Quantity, operation.Type)
		s.CalculateWeightedAvgPrice(i)
		tax := operation.CalculateTaxes(s)
		taxes = append(taxes, tax)
	}

	return taxes
}

func NewStock(operations []Operation) Stock {
	stock := Stock{
		operations: operations,
	}

	return stock
}
