package main

import "testing"

func TestAdd1(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := addFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 5 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "5_0" {
		t.Fatalf("Expected 5_0. Actual " + res.ToString() + "\n")
	}
}

func TestAdd2(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}
	res, err := addFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 1 || res.numerator != 0 || res.denominator != 0 || res.ToString() != "1_0" {
		t.Fatalf("Expected 1_0. Actual " + res.ToString() + "\n")
	}
}

func TestAdd3(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 23, numerator: 33, denominator: 67}
	rhs := &fraction{isNegative: false, whole: 56, numerator: 45, denominator: 839}
	res, err := addFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 79 || res.numerator != 30702 || res.denominator != 56213 || res.ToString() != "79_30702/56213" {
		t.Fatalf("Expected 79_30702/56213. Actual " + res.ToString() + "\n")
	}
}

func TestSubtract(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 1, numerator: 3, denominator: 6}
	res, err := subtractFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 1 || res.numerator != 0 || res.denominator != 0 || res.ToString() != "1_0" {
		t.Fatalf("Expected 1_0. Actual " + res.ToString() + "\n")
	}
}

func TestSubtractToZero(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := subtractFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 0 || res.numerator != 0 || res.denominator != 6 || res.ToString() != "0" {
		t.Fatalf("Expected 0. Actual " + res.ToString() + "\n")
	}
}

func TestSubtract2(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 7, numerator: 5, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 3, numerator: 2, denominator: 8}
	res, err := subtractFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 4 || res.numerator != 7 || res.denominator != 12 || res.ToString() != "4_7/12" {
		t.Fatalf("Expected 4_7/12. Actual " + res.ToString() + "\n")
	}
}

func TestSubtractNegative(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 1, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := subtractFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 1 || res.numerator != 0 || res.denominator != 0 || res.ToString() != "-1_0" {
		t.Fatalf("Expected -1_0. Actual " + res.ToString() + "\n")
	}
}

func TestSubtractNegative2(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 3, numerator: 2, denominator: 8}
	rhs := &fraction{isNegative: false, whole: 7, numerator: 5, denominator: 6}
	res, err := subtractFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 4 || res.numerator != 7 || res.denominator != 12 || res.ToString() != "-4_7/12" {
		t.Fatalf("Expected -4_7/12. Actual " + res.ToString() + "\n")
	}
}

func TestMultiply(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 3, numerator: 2, denominator: 8}
	rhs := &fraction{isNegative: false, whole: 7, numerator: 5, denominator: 6}
	res, err := multiplyFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 25 || res.numerator != 11 || res.denominator != 24 || res.ToString() != "25_11/24" {
		t.Fatalf("Expected 25_11/24. Actual " + res.ToString() + "\n")
	}
}

func TestMultiplyZero(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 0, denominator: 1}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 0, denominator: 1}
	res, err := multiplyFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 0 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "0" {
		t.Fatalf("Expected 0. Actual " + res.ToString() + "\n")
	}
}

func TestDivide(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 3, numerator: 2, denominator: 8}
	rhs := &fraction{isNegative: false, whole: 7, numerator: 5, denominator: 6}
	res, err := divideFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 0 || res.numerator != 39 || res.denominator != 94 || res.ToString() != "39/94" {
		t.Fatalf("Expected 39/94. Actual " + res.ToString() + "\n")
	}
}

func TestDivideZero(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 0, denominator: 1}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 0, denominator: 1}
	res, err := divideFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 0 || res.numerator != 0 || res.denominator != 0 || res.ToString() != "0" {
		t.Fatalf("Expected 0. Actual " + res.ToString() + "\n")
	}
}
