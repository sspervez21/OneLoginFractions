package main

import "testing"

// Both operands positive
func TestSignedAdd1(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := addSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 8 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "8_0" {
		t.Fatalf("Expected 8_0. Actual " + res.ToString() + "\n")
	}
}

// lhs negative rhs positive
func TestSignedAdd2(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := addSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 3 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "-3_0" {
		t.Fatalf("Expected -3_0. Actual " + res.ToString() + "\n")
	}
}

// rhs negative lhs positive
func TestSignedAdd3(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := addSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 3 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "3_0" {
		t.Fatalf("Expected 3_0. Actual " + res.ToString() + "\n")
	}
}

// Both operands negative
func TestSignedAdd4(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := addSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 8 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "-8_0" {
		t.Fatalf("Expected -8_0. Actual " + res.ToString() + "\n")
	}
}

// Both operands positive
func TestSignedSubtract1(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := subtractSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 3 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "3_0" {
		t.Fatalf("Expected 3_0. Actual " + res.ToString() + "\n")
	}
}

// lhs negative rhs positive
func TestSignedSubtract2(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := subtractSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 8 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "-8_0" {
		t.Fatalf("Expected -8_0. Actual " + res.ToString() + "\n")
	}
}

// rhs negative lhs positive
func TestSignedSubtract3(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := subtractSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 8 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "8_0" {
		t.Fatalf("Expected 8_0. Actual " + res.ToString() + "\n")
	}
}

// Both operands negative
func TestSignedSubtract4(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := subtractSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 3 || res.numerator != 0 || res.denominator != 1 || res.ToString() != "-3_0" {
		t.Fatalf("Expected -3_0. Actual " + res.ToString() + "\n")
	}
}

// Both operands positive
func TestSignedMultiply1(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := multiplySignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 13 || res.numerator != 3 || res.denominator != 4 || res.ToString() != "13_3/4" {
		t.Fatalf("Expected 13_3/4. Actual " + res.ToString() + "\n")
	}
}

// lhs negative rhs positive
func TestSignedMultiply2(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := multiplySignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 13 || res.numerator != 3 || res.denominator != 4 || res.ToString() != "-13_3/4" {
		t.Fatalf("Expected -13_3/4. Actual " + res.ToString() + "\n")
	}
}

// rhs negative lhs positive
func TestSignedMultiply3(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := multiplySignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 13 || res.numerator != 3 || res.denominator != 4 || res.ToString() != "-13_3/4" {
		t.Fatalf("Expected -13_3/4. Actual " + res.ToString() + "\n")
	}
}

// Both operands negative
func TestSignedMultiply4(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := multiplySignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 13 || res.numerator != 3 || res.denominator != 4 || res.ToString() != "13_3/4" {
		t.Fatalf("Expected 13_3/4. Actual " + res.ToString() + "\n")
	}
}

// Both operands positive
func TestSignedDivide1(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := divideSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 2 || res.numerator != 1 || res.denominator != 5 || res.ToString() != "2_1/5" {
		t.Fatalf("Expected 2_1/5. Actual " + res.ToString() + "\n")
	}
}

// lhs negative rhs positive
func TestSignedDivide2(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 2, numerator: 3, denominator: 6}
	res, err := divideSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 2 || res.numerator != 1 || res.denominator != 5 || res.ToString() != "-2_1/5" {
		t.Fatalf("Expected -2_1/5. Actual " + res.ToString() + "\n")
	}
}

// rhs negative lhs positive
func TestSignedDivide3(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := divideSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if !res.isNegative || res.whole != 2 || res.numerator != 1 || res.denominator != 5 || res.ToString() != "-2_1/5" {
		t.Fatalf("Expected -2_1/5. Actual " + res.ToString() + "\n")
	}
}

// Both operands negative
func TestSignedDivide4(t *testing.T) {
	lhs := &fraction{isNegative: true, whole: 5, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: true, whole: 2, numerator: 3, denominator: 6}
	res, err := divideSignedFractions(lhs, rhs)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error() + "\n")
	}
	if res.isNegative || res.whole != 2 || res.numerator != 1 || res.denominator != 5 || res.ToString() != "2_1/5" {
		t.Fatalf("Expected 2_1/5. Actual " + res.ToString() + "\n")
	}
}
