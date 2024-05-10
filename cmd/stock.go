package main

type Stock struct {
	weightedAvgPrice float32
	operations       []Operation
	totalQuantity    int32
	profit           float32
	losses           float32
}

func (s *Stock) UpdateTotalAmount(quantity int32, operationType string) {
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

func (s *Stock) CalculateAvgPrice(currentQuantity int32, operationQuantity int32, operationUnitCost float32) float32 {
	avgPrice := ((float32(currentQuantity) * s.weightedAvgPrice) +
		(float32(operationQuantity) * operationUnitCost)) /
		float32(currentQuantity+operationQuantity)

	s.weightedAvgPrice = avgPrice
	return avgPrice
}

func (s *Stock) GetTaxes() []map[string]float32 {
	taxes := make([]map[string]float32, 0, 0)
	var currentQuantity int32 = 0
	for i, operation := range s.operations {
		isFirstBuy := i == 0
		isSellOperation := operation.Type == SellOperation
		isBuyOperation := operation.Type == BuyOperation

		if isFirstBuy {
			currentQuantity = operation.Quantity
			s.weightedAvgPrice = operation.UnitCost
		}

		if isSellOperation {
			currentQuantity = currentQuantity - operation.Quantity
		}

		if isBuyOperation && !isFirstBuy {
			s.CalculateAvgPrice(currentQuantity, operation.Quantity, operation.UnitCost)
			currentQuantity = currentQuantity + operation.Quantity
		}

		s.UpdateTotalAmount(operation.Quantity, operation.Type)

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
