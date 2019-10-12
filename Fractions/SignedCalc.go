package main

import "errors"

func addSignedFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	if !lhs.isNegative && !rhs.isNegative {
		return addFractions(lhs, rhs)
	} else if !lhs.isNegative && rhs.isNegative {
		rhs.isNegative = false
		return subtractFractions(lhs, rhs)
	} else if lhs.isNegative && !rhs.isNegative {
		lhs.isNegative = false
		return subtractFractions(rhs, lhs)
	} else if lhs.isNegative && rhs.isNegative {
		lhs.isNegative = false
		rhs.isNegative = false
		res, err := addFractions(lhs, rhs)

		if err == nil {
			res.isNegative = true
		}

		return res, err
	} else {
		return nil, errors.New("unknown sign combination")
	}
}

func subtractSignedFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	if !lhs.isNegative && !rhs.isNegative {
		return subtractFractions(lhs, rhs)
	} else if !lhs.isNegative && rhs.isNegative {
		rhs.isNegative = false
		return addFractions(lhs, rhs)
	} else if lhs.isNegative && !rhs.isNegative {
		lhs.isNegative = false
		res, err := addFractions(lhs, rhs)

		if err == nil {
			res.isNegative = true
		}

		return res, err
	} else if lhs.isNegative && rhs.isNegative {
		lhs.isNegative = false
		rhs.isNegative = false
		return subtractFractions(rhs, lhs)
	} else {
		return nil, errors.New("unknown sign combination")
	}
}

func multiplySignedFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	if !lhs.isNegative && !rhs.isNegative {
		return multiplyFractions(lhs, rhs)
	} else if !lhs.isNegative && rhs.isNegative {
		rhs.isNegative = false
		res, err := multiplyFractions(lhs, rhs)

		if err == nil {
			res.isNegative = true
		}

		return res, err
	} else if lhs.isNegative && !rhs.isNegative {
		lhs.isNegative = false
		res, err := multiplyFractions(lhs, rhs)

		if err == nil {
			res.isNegative = true
		}

		return res, err
	} else if lhs.isNegative && rhs.isNegative {
		lhs.isNegative = false
		rhs.isNegative = false
		return multiplyFractions(lhs, rhs)
	} else {
		return nil, errors.New("unknown sign combination")
	}
}

func divideSignedFractions(lhs *fraction, rhs *fraction) (*fraction, error) {
	if !lhs.isNegative && !rhs.isNegative {
		return divideFractions(lhs, rhs)
	} else if !lhs.isNegative && rhs.isNegative {
		rhs.isNegative = false
		res, err := divideFractions(lhs, rhs)

		if err == nil {
			res.isNegative = true
		}

		return res, err
	} else if lhs.isNegative && !rhs.isNegative {
		lhs.isNegative = false
		res, err := divideFractions(lhs, rhs)

		if err == nil {
			res.isNegative = true
		}

		return res, err
	} else if lhs.isNegative && rhs.isNegative {
		lhs.isNegative = false
		rhs.isNegative = false
		return divideFractions(lhs, rhs)
	} else {
		return nil, errors.New("unknown sign combination")
	}
}
