package main

import (
	"calculator"
	"fmt"
)

func main() {

	addition := calculator.AddMany(2, 2)
	fmt.Println(" - Addition result: \t\t", addition)

	substract := calculator.SubstractMany(5, 1, 2)
	fmt.Println(" - Substraction result: \t", substract)

	multiply := calculator.MultiplyMany(5, 1, 2)
	fmt.Println(" - Multiplication result: \t", multiply)

	divide, err := calculator.DivideMany(5, 1, 2)
	if err == nil {
		fmt.Println(" - Division result: \t\t", divide)
	}
}
