package main

import (
	"fmt"
)

type CalcNumber struct {
	value int
}

func CreateCalculableNumber (value int) CalcNumber {
	return CalcNumber { value }
}

func (number CalcNumber) sum (valueToSum int) int {
	return number.value + valueToSum
}

func (number CalcNumber) subtract (valueToSubtract int) int {
	return number.value - valueToSubtract
}

func (number CalcNumber) multiplyBy (valueToMultiplyBy int) int {
	return number.value * valueToMultiplyBy
}

func (number CalcNumber) divideBy (valueToDivideBy int) int {
	return number.value / valueToDivideBy
}

func main () {
	numbersToCalculate := []int {10, 20, 30, 40, 50, 60}

	for _, number := range numbersToCalculate {
		calculableNumber := CreateCalculableNumber(number)

		algo := 2

		fmt.Printf("%d\n\n", calculableNumber.value);

    fmt.Printf("%d + %d = %d\n", calculableNumber.value, algo, calculableNumber.sum(algo));
    fmt.Printf("%d - %d = %d\n", calculableNumber.value, algo, calculableNumber.subtract(algo));
    fmt.Printf("%d * %d = %d\n", calculableNumber.value, algo, calculableNumber.multiplyBy(algo));
    fmt.Printf("%d / %d = %d\n", calculableNumber.value, algo, calculableNumber.divideBy(algo));
    
    fmt.Printf("\n");
	}
}