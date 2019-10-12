package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Operator enum
type Operator int

const (
	plus     Operator = 0
	minus    Operator = 1
	multiply Operator = 2
	divide   Operator = 3
	unknown  Operator = 4
)

type fraction struct {
	isNegative  bool
	whole       uint64
	numerator   uint64
	denominator uint64
}

func parseOperator(op string) (Operator, error) {
	if op == "+" {
		return plus, nil
	}
	if op == "-" {
		return minus, nil
	}
	if op == "*" {
		return multiply, nil
	}
	if op == "/" {
		return divide, nil
	}
	return unknown, errors.New("could not read operator. allowed operators are +,-,*,/")
}

func parseFraction(s string) (*fraction, error) {

	// TODO: Parse whole number only...
	// TODO: Break this function up!

	fr := s
	if len(fr) == 0 {
		return nil, errors.New("could not parse empty fraction")
	}

	fraction := &fraction{}

	// parse negative number
	if strings.HasPrefix(fr, "-") {
		fraction.isNegative = true

		// remove optional '-' sign
		fr = fr[1:]

		if len(fr) == 0 {
			return nil, errors.New("could not parse fraction with no number component:" + s)
		}
	}

	// parse mixed fraction
	if strings.Contains(fr, "_") {
		idx := strings.Index(fr, "_")
		whole := fr[:idx]

		u, err := strconv.ParseUint(whole, 10, 64)

		if err != nil {
			return nil, errors.New("could not parse whole number component:" + s)
		}

		fraction.whole = u

		if len(fr) == idx+1 {
			return nil, errors.New("could not parse fraction, missing fractional component:" + s)
		}

		fr = fr[idx+1:]
	}

	if !strings.Contains(fr, "/") {
		return nil, errors.New("could not parse fraction, fractional component must be specified as numerator/denominator:" + s)
	}

	idx := strings.Index(fr, "/")

	// there should be something before and after the /
	if idx == 0 || idx == (len(fr)-1) {
		return nil, errors.New("could not parse fraction, fractional component must have both a numerator and denominator:" + s)
	}

	numeratorStr := fr[:idx]
	denominatorStr := fr[idx+1:]

	numerator, err := strconv.ParseUint(numeratorStr, 10, 64)
	if err != nil {
		return nil, errors.New("could not parse numerator:" + s)
	}

	denominator, err := strconv.ParseUint(denominatorStr, 10, 64)
	if err != nil {
		return nil, errors.New("could not parse denominator:" + s)
	}

	fraction.numerator = numerator
	fraction.denominator = denominator

	if fraction.denominator == 0 {
		return nil, errors.New("a denominator of 0 is not allowed")
	}

	return fraction, nil
}

func (fr *fraction) ToString() string {
	ret := ""

	if fr.numerator == 0 {
		ret += "0"
	} else {
		ret += fmt.Sprint(fr.numerator) + "/" + fmt.Sprint(fr.denominator)
	}

	if fr.whole != 0 {
		ret = fmt.Sprint(fr.whole) + "_" + ret
	}

	if fr.isNegative {
		ret = "-" + ret
	}

	return ret
}

func convertMixedFractionToImproperFraction(fr *fraction) (*fraction, error) {
	if fr.denominator == 0 {
		return nil, errors.New("expected mixed fraction to have a non-zero denominator." + fr.ToString())
	}

	if fr.whole == 0 {
		return fr, nil
	}

	// TODO: check for int overflow
	fr.numerator = (fr.whole * fr.denominator) + fr.numerator
	fr.whole = 0
	return fr, nil
}

func findCommonDenominator(lhs *fraction, rhs *fraction) (uint64, error) {
	if lhs.whole != 0 || rhs.whole != 0 {
		return 0, errors.New("expect proper fractions")
	}

	if lhs.denominator == rhs.denominator {
		return lhs.denominator, nil
	}

	if lhs.denominator < rhs.denominator && rhs.denominator%lhs.denominator == 0 {
		return rhs.denominator, nil
	}

	if lhs.denominator > rhs.denominator && lhs.denominator%rhs.denominator == 0 {
		return lhs.denominator, nil
	}

	// TODO: check for overflow
	// TODO: rewrite this function to use LCD (perhaps using prime factors?). Need to validate performance tradeoffs
	return lhs.denominator * rhs.denominator, nil
}

func (fr *fraction) simplify() {
	// TODO: should we still simplify? Maybe...
	if fr.whole != 0 {
		return
	}

	// cannot simplify further
	if fr.numerator == 0 || fr.denominator == 0 {
		return
	}

	if fr.numerator == fr.denominator {
		fr.whole++
		fr.numerator = 0
		fr.denominator = 0
	}

	gcd := findGCD(fr.numerator, fr.denominator)

	if gcd > 1 {
		fr.numerator = fr.numerator / gcd
		fr.denominator = fr.denominator / gcd
	}

	// TODO: what if numerator == denominator?
	// This should not be the case but we should add a debug/assert to validate this perhaps.
	if fr.numerator > fr.denominator {
		fr.whole = fr.numerator / fr.denominator
		fr.numerator = fr.numerator % fr.denominator
	}
}

func findGCD(a uint64, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
