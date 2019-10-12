package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Printf("? ")

		inputStr, err := readInput()
		if err != nil {
			fmt.Printf("Error::" + err.Error() + ".\n")
			continue
		}

		inputStr = strings.TrimSpace(inputStr)

		if strings.EqualFold(inputStr, "exit") || strings.EqualFold(inputStr, "quit") {
			break
		}

		ops := strings.Fields(inputStr)

		if len(ops) != 3 {
			fmt.Print("Error::Unexpected input. Input should follow the pattern \"1/2 * 3_3/4\".\n")
			continue
		}

		lhs, err := parseFraction(ops[0])
		if err != nil {
			fmt.Print("Error::" + err.Error() + ".\n")
			continue
		}

		op, err := parseOperator(ops[1])
		if err != nil {
			fmt.Print("Error::" + err.Error() + ".\n")
			continue
		}

		rhs, err := parseFraction(ops[2])
		if err != nil {
			fmt.Print("Error::" + err.Error() + ".\n")
			continue
		}

		res, err := calculateResult(lhs, op, rhs)
		if err != nil {
			fmt.Print("Error::" + err.Error() + ".\n")
			continue
		}

		fmt.Print("= " + res.ToString() + "\n")
	}
}

func calculateResult(lhs *fraction, op Operator, rhs *fraction) (*fraction, error) {
	if op == plus {
		return addSignedFractions(lhs, rhs)
	} else if op == minus {
		return subtractSignedFractions(lhs, rhs)
	} else if op == multiply {
		return multiplySignedFractions(lhs, rhs)
	} else if op == divide {
		return divideSignedFractions(lhs, rhs)
	} else {
		return nil, errors.New("unknow operator, could not perform calculation")
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	// We may ignore this error. It is returned if the reader is unable to find the delimiter which is unexpected behavior.
	text, err := reader.ReadString('\n')
	return text, err
}
