package main

import (
	"fmt"
	"strings"
)

const (
	LimitError = "[!] число должно быть в диапазоне от 1 до 10 включительно"
	NumberSystemsError = "[!] Используются одновременно разные системы счисления"
	NotSupportedEquation = "[!] Формат математической операции не удовлетворяет заданию два операнда и один оператор (+, -, /, *)"
)

func checkLimit(num int) error {
	if num <= 10 && num >= 1 {
		return nil
	}

	return fmt.Errorf(LimitError)
}


func parseInput(input string) (Equation, NumberType, error) {
	parts := strings.Split(input, " ")

	if  len(parts) < 3{
		return Equation{}, 0,fmt.Errorf(IncorrectEquationError)
	}

	if len(parts) > 3 || len(parts[1]) != 1 {
		return Equation{}, 0,fmt.Errorf(NotSupportedEquation)
	}

	num1, num1Type, err1 := convertNumber(parts[0])
	num2, num2Type, err2 := convertNumber(parts[2])

	if (err1 != nil){
		return Equation{}, 0,err1
	}

	if (err2 != nil){
		return Equation{}, 0,err2
	}

	if (num1Type != num2Type){
		return Equation{}, 0,fmt.Errorf(NumberSystemsError)
	}

	err := checkLimit(num1)
	if err != nil {
    	return Equation{}, 0,err 
	}

	err = checkLimit(num2) // исправление
	if err != nil {
    	return Equation{}, 0,err 
	}

	op := []byte(parts[1])[0]
	err = opCheck(op)

	if (err != nil){
		return Equation{}, 0,err
	}

	return Equation{a: num1, b: num2, op: op}, num1Type, nil
}

func opCheck(op byte) error {
	for _, opChecking := range SupportedOperations {
		if opChecking == op{
			return nil
		}
	}

	return fmt.Errorf(IncorrectEquationError)
}