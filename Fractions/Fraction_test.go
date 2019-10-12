package main

import (
	"strconv"
	"testing"
)

func TestParseOperator(t *testing.T) {

	op, err := parseOperator("+")
	if op != plus || err != nil {
		t.Fatalf("Operator + should be parsed as Operator.plus.\n")
	}

	op, err = parseOperator("-")
	if op != minus || err != nil {
		t.Fatalf("Operator - should be parsed as Operator.minus.\n")
	}

	op, err = parseOperator("*")
	if op != multiply || err != nil {
		t.Fatalf("Operator * should be parsed as Operator.multiply.\n")
	}

	op, err = parseOperator("/")
	if op != divide || err != nil {
		t.Fatalf("Operator / should be parsed as Operator.divide.\n")
	}

	op, err = parseOperator("")
	if err == nil {
		t.Fatalf("Parsing empty operator should return error.\n")
	}

	op, err = parseOperator("abc")
	if err == nil {
		t.Fatalf("Parsing incorrect operator should return error.\n")
	}
}

func TestParseFraction(t *testing.T) {

	fr, err := parseFraction("334/4567")

	if err != nil {
		t.Fatalf("Unexpected error while parsing mixed fraction.")
	}

	if fr.isNegative {
		t.Fatalf("Non-negative fraction expected.")
	}

	if fr.numerator != 334 {
		t.Fatalf("Incorrect numerator.")
	}

	if fr.denominator != 4567 {
		t.Fatalf("Incorrect denominator.")
	}
}

func TestParseImproperFraction(t *testing.T) {

	fr, err := parseFraction("3345/56")

	if err != nil {
		t.Fatalf("Unexpected error while parsing mixed fraction.")
	}

	if fr.isNegative {
		t.Fatalf("Non-negative fraction expected.")
	}

	if fr.numerator != 3345 {
		t.Fatalf("Incorrect numerator.")
	}

	if fr.denominator != 56 {
		t.Fatalf("Incorrect denominator.")
	}
}

func TestParseMixedFraction(t *testing.T) {

	fr, err := parseFraction("2_3/4")

	if err != nil {
		t.Fatalf("Unexpected error while parsing mixed fraction.")
	}

	if fr.isNegative {
		t.Fatalf("Non-negative fraction expected.")
	}

	if fr.whole != 2 {
		t.Fatalf("Incorrect whole number.")
	}

	if fr.numerator != 3 {
		t.Fatalf("Incorrect numerator.")
	}

	if fr.denominator != 4 {
		t.Fatalf("Incorrect denominator.")
	}
}
func TestParseMixedImproperFraction(t *testing.T) {

	fr, err := parseFraction("256_36894/434")

	if err != nil {
		t.Fatalf("Unexpected error while parsing mixed fraction.")
	}

	if fr.isNegative {
		t.Fatalf("Non-negative fraction expected.")
	}

	if fr.whole != 256 {
		t.Fatalf("Incorrect whole number.")
	}

	if fr.numerator != 36894 {
		t.Fatalf("Incorrect numerator.")
	}

	if fr.denominator != 434 {
		t.Fatalf("Incorrect denominator.")
	}
}

func TestToString(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 0, numerator: 15, denominator: 256}
	fs := fr.ToString()
	if fs != "15/256" {
		t.Fatalf("Unexpected conversion to string. Expected: 15/256 Actual: " + fs)
	}
}

func TestToStringImproperFraction(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 6, numerator: 15, denominator: 256}
	fs := fr.ToString()
	if fs != "6_15/256" {
		t.Fatalf("Unexpected conversion to string. Expected: 6_15/256 actual: " + fs)
	}
}

func TestToStringNegative(t *testing.T) {
	fr := &fraction{isNegative: true, whole: 0, numerator: 15, denominator: 256}
	fs := fr.ToString()
	if fs != "-15/256" {
		t.Fatalf("Unexpected conversion to string. Expected: -15/256 actual: " + fs)
	}
}
func TestToStringImproperFractionNegative(t *testing.T) {
	fr := &fraction{isNegative: true, whole: 6, numerator: 15, denominator: 256}
	fs := fr.ToString()
	if fs != "-6_15/256" {
		t.Fatalf("Unexpected conversion to string. Expected: -6_15/256 actual: " + fs)
	}
}
func TestConvertMixedFractionToImproperFraction(t *testing.T) {
	mixedfr := &fraction{isNegative: false, whole: 6, numerator: 15, denominator: 256}
	improperfr, err := convertMixedFractionToImproperFraction(mixedfr)

	if err != nil {
		t.Fatalf("Unexpected error." + err.Error())
	}

	if improperfr.isNegative || improperfr.whole != 0 || improperfr.numerator != 1551 || improperfr.denominator != 256 {
		t.Fatalf("Unexpected result. Expected 1551/256 actual " + improperfr.ToString())
	}
}
func TestFindCommonDenominator1(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 1, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}

	_, err := findCommonDenominator(lhs, rhs)

	if err == nil {
		t.Fatalf("Expected error. We only simplify fractions with no whole number component.")
	}
}
func TestFindCommonDenominator2(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 1, numerator: 3, denominator: 6}

	_, err := findCommonDenominator(lhs, rhs)

	if err == nil {
		t.Fatalf("Expected error. We only simplify fractions with no whole number component.")
	}
}
func TestFindCommonDenominator3(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}

	comDenom, err := findCommonDenominator(lhs, rhs)

	if err != nil {
		t.Fatalf("Unexpected error. " + err.Error())
	}

	if comDenom != 6 {
		t.Fatalf("Expected 6 Actual " + strconv.FormatUint(comDenom, 10))
	}
}
func TestFindCommonDenominator4(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 18}

	comDenom, err := findCommonDenominator(lhs, rhs)

	if err != nil {
		t.Fatalf("Unexpected error. " + err.Error())
	}

	if comDenom != 18 {
		t.Fatalf("Expected 18 Actual " + strconv.FormatUint(comDenom, 10))
	}
}
func TestFindCommonDenominator5(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 18}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 6}

	comDenom, err := findCommonDenominator(lhs, rhs)

	if err != nil {
		t.Fatalf("Unexpected error. " + err.Error())
	}

	if comDenom != 18 {
		t.Fatalf("Expected 18 Actual " + strconv.FormatUint(comDenom, 10))
	}
}
func TestFindCommonDenominator6(t *testing.T) {
	lhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 7}
	rhs := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 45}

	comDenom, err := findCommonDenominator(lhs, rhs)

	if err != nil {
		t.Fatalf("Unexpected error. " + err.Error())
	}

	if comDenom != 315 {
		t.Fatalf("Expected 315 Actual " + strconv.FormatUint(comDenom, 10))
	}
}
func TestSimplify1(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 3, numerator: 3, denominator: 7}

	fr.simplify()

	if fr.isNegative || fr.whole != 3 || fr.numerator != 3 || fr.denominator != 7 {
		t.Fatalf("Expected fraction to remain unchanged." + fr.ToString())
	}
}
func TestSimplify2(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 0, numerator: 3, denominator: 7}

	fr.simplify()

	if fr.isNegative || fr.whole != 0 || fr.numerator != 3 || fr.denominator != 7 {
		t.Fatalf("Expected fraction to remain unchanged." + fr.ToString())
	}
}
func TestSimplify3(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 0, numerator: 9, denominator: 21}

	fr.simplify()

	if fr.isNegative || fr.whole != 0 || fr.numerator != 3 || fr.denominator != 7 {
		t.Fatalf("Expected fraction 0_3/7 Actual " + fr.ToString())
	}
}
func TestSimplify4(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 0, numerator: 0, denominator: 21}

	fr.simplify()

	if fr.isNegative || fr.whole != 0 || fr.numerator != 0 || fr.denominator != 21 {
		t.Fatalf("Expected fraction to remain unchanged." + fr.ToString())
	}
}
func TestSimplify5(t *testing.T) {
	fr := &fraction{isNegative: false, whole: 0, numerator: 9, denominator: 0}

	fr.simplify()

	if fr.isNegative || fr.whole != 0 || fr.numerator != 9 || fr.denominator != 0 {
		t.Fatalf("Expected fraction to remain unchanged." + fr.ToString())
	}
}
