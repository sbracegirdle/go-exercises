package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Operation struct {
	Func       func(int, int) int
	Precedence int
	Name       string
}

var Add = Operation{func(a, b int) int { return a + b }, 1, "Add"}
var Subtract = Operation{func(a, b int) int { return a - b }, 1, "Subtract"}
var Multiply = Operation{func(a, b int) int { return a * b }, 2, "Multiply"}
var Divide = Operation{func(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}, 2, "Divide"}
var Exponent = Operation{func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }, 3, "Exponent"}

func Parse(expression string) ([]int, []Operation, error) {
	numbers := []int{}
	operations := []Operation{}
	numStr := ""

	for _, char := range expression {
		if char >= '0' && char <= '9' {
			numStr += string(char)
		} else if strings.Contains("+-*/^", string(char)) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing part %q as integer: %w", numStr, err)
			}
			numbers = append(numbers, num)
			numStr = ""

			switch char {
			case '+':
				operations = append(operations, Add)
			case '-':
				operations = append(operations, Subtract)
			case '*':
				operations = append(operations, Multiply)
			case '/':
				operations = append(operations, Divide)
			case '^':
				operations = append(operations, Exponent)
			}
		} else if char != ' ' {
			return nil, nil, fmt.Errorf("unknown character: %c", char)
		}
	}

	if numStr != "" {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing part %q as integer: %w", numStr, err)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) != len(operations)+1 {
		return nil, nil, fmt.Errorf("mismatch between numbers and operations")
	}

	return numbers, operations, nil
}

func Calculate(numbers []int, operations []Operation) (int, error) {
	for precedence := 3; precedence > 0; precedence-- {
		for i := len(operations) - 1; i >= 0; i-- {
			operation := operations[i]
			if operation.Precedence == precedence {
				if len(numbers) != len(operations)+1 {
					return 0, fmt.Errorf("mismatch between numbers and operations, numbers: %v, operations: %v", numbers, operations)
				}
				// check i is not beyond the end of the slices
				if i+1 >= len(numbers) || i >= len(operations) {
					return 0, fmt.Errorf("index out of range, i: %d, numbers: %v, operations: %v", i, numbers, operations)
				}

				numbers[i] = operation.Func(numbers[i], numbers[i+1])
				numbers = append(numbers[:i+1], numbers[i+2:]...)
				operations = append(operations[:i], operations[i+1:]...)
				i-- // adjust index for removed operation
			}
		}
	}

	// After all operations with higher precedence are done, we should be left with only addition and subtraction
	result := numbers[0]
	for i, operation := range operations {
		result = operation.Func(result, numbers[i+1])
	}

	return result, nil
}

func main() {
	expression := "1 + 2 * 3"
	numbers, operations, err := Parse(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err := Calculate(numbers, operations)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Result:", result)
	}
}
