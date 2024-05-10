package tests

import (
	. "github.com/enkdress/capital-gains"
	"reflect"
	"testing"
)

var case7Operations = []Operation{
	{Type: "buy", UnitCost: 10, Quantity: 10000},
	{Type: "sell", UnitCost: 2, Quantity: 5000},
	{Type: "sell", UnitCost: 20, Quantity: 2000},
	{Type: "sell", UnitCost: 20, Quantity: 2000},
	{Type: "sell", UnitCost: 25, Quantity: 1000},
	{Type: "buy", UnitCost: 20, Quantity: 10000},
	{Type: "sell", UnitCost: 15, Quantity: 5000},
	{Type: "sell", UnitCost: 30, Quantity: 4350},
	{Type: "sell", UnitCost: 30, Quantity: 650},
}

func TestNewStock(t *testing.T) {
	type args struct {
		operations []Operation
	}
	tests := []struct {
		name string
		args args
		want Stock
	}{
		{
			name: "Test Case 7",
			args: args{
				operations: case7Operations,
			},
			want: Stock{
				weightedAvgPrice: 20,
				operations:       case7Operations,
				profit:           0,
				taxes: []map[string]float32{
					{"tax": 0.00}, {"tax": 0.00}, {"tax": 0.00}, {"tax": 0.00}, {"tax": 3000.00},
					{"tax": 0.00}, {"tax": 0.00}, {"tax": 3700.00}, {"tax": 0.00},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStock(tt.args.operations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStock() = %v, want %v", got, tt.want)
			}
		})
	}
}
