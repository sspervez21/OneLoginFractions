package main

func addFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	lhs, rhs, err := normalize(lhs, rhs)
	if err != nil {
		return nil, err
	}

	result := &fraction{isNegative: false, whole: 0, numerator: lhs.numerator + rhs.numerator, denominator: lhs.denominator}
	result.simplify()

	return result, nil
}

func subtractFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	lhs, rhs, err := normalize(lhs, rhs)
	if err != nil {
		return nil, err
	}

	result := &fraction{isNegative: false, whole: 0, numerator: lhs.numerator - rhs.numerator, denominator: lhs.denominator}
	if lhs.numerator < rhs.numerator {
		result.numerator = rhs.numerator - lhs.numerator
		result.isNegative = true
	}

	result.simplify()
	return result, nil
}

func multiplyFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	lhs, err := convertMixedFractionToImproperFraction(lhs)
	if err != nil {
		return nil, err
	}

	rhs, err = convertMixedFractionToImproperFraction(rhs)
	if err != nil {
		return nil, err
	}

	result := &fraction{isNegative: false, whole: 0, numerator: lhs.numerator * rhs.numerator, denominator: lhs.denominator * rhs.denominator}

	result.simplify()
	return result, nil
}

func divideFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	lhs, err := convertMixedFractionToImproperFraction(lhs)
	if err != nil {
		return nil, err
	}

	rhs, err = convertMixedFractionToImproperFraction(rhs)
	if err != nil {
		return nil, err
	}

	result := &fraction{isNegative: false, whole: 0, numerator: lhs.numerator * rhs.denominator, denominator: lhs.denominator * rhs.numerator}

	result.simplify()
	return result, nil
}

func normalize(lhs *fraction, rhs *fraction) (*fraction, *fraction, error) {
	lhs, err := convertMixedFractionToImproperFraction(lhs)
	if err != nil {
		return nil, nil, err
	}

	rhs, err = convertMixedFractionToImproperFraction(rhs)
	if err != nil {
		return nil, nil, err
	}

	commonDenominator, err := findCommonDenominator(lhs, rhs)
	if err != nil {
		return nil, nil, err
	}

	lhs.numerator = (commonDenominator / lhs.denominator) * lhs.numerator
	lhs.denominator = (commonDenominator / lhs.denominator) * lhs.denominator

	rhs.numerator = (commonDenominator / rhs.denominator) * rhs.numerator
	rhs.denominator = (commonDenominator / rhs.denominator) * rhs.denominator

	return lhs, rhs, nil
}
