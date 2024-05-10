package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case1Input.taxes, Case1Result)
	if !isTaxEqual {
		t.Errorf("\nCase 1 failed:\nExpected: %v\nRecieved: %v", Case1Result, Case1Input.taxes)
	}
}

func TestCase2(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case2Input.taxes, Case2Result)
	if !isTaxEqual {
		t.Errorf("\nCase 2 failed:\nExpected: %v\nRecieved: %v", Case2Result, Case2Input.taxes)
	}
}

func TestCase3(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case3Input.taxes, Case3Result)
	if !isTaxEqual {
		t.Errorf("\nCase 3 failed:\nExpected: %v\nRecieved: %v", Case3Result, Case3Input.taxes)
	}
}

func TestCase4(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case4Input.taxes, Case4Result)
	if !isTaxEqual {
		t.Errorf("\nCase 4 failed:\nExpected: %v\nRecieved: %v", Case4Result, Case4Input.taxes)
	}
}

func TestCase5(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case5Input.taxes, Case5Result)
	if !isTaxEqual {
		t.Errorf("\nCase 5 failed:\nExpected: %v\nRecieved: %v", Case5Result, Case5Input.taxes)
	}
}

func TestCase6(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case6Input.taxes, Case6Result)
	if !isTaxEqual {
		t.Errorf("\nCase 6 failed:\nExpected: %v\nRecieved: %v", Case6Result, Case6Input.taxes)
	}
}

func TestCase7(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case7Input.taxes, Case7Result)
	if !isTaxEqual {
		t.Errorf("\nCase 7 failed:\nExpected: %v\nRecieved: %v", Case7Result, Case7Input.taxes)
	}
}

func TestCase8(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case8Input.taxes, Case8Result)
	if !isTaxEqual {
		t.Errorf("\nCase 8 failed:\nExpected: %v\nRecieved: %v", Case8Result, Case8Input.taxes)
	}
}

func TestCase9(t *testing.T) {
	isTaxEqual := reflect.DeepEqual(Case9Input.taxes, Case9Result)
	if !isTaxEqual {
		t.Errorf("\nCase 9 failed:\nExpected: %v\nRecieved: %v", Case9Result, Case9Input.taxes)
	}
}

