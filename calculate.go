package main

import "fmt"

const (
	DivisionByZeroError = "[!] Деление на ноль"
)

var SupportedOperations = [...]byte {'+', '-', '*', '/'}

func calculate(eq Equation) (int, error) {
	switch eq.op {
	case '+': return eq.a + eq.b, nil
	case '-': return eq.a - eq.b, nil
	case '*': return eq.a * eq.b, nil
	case '/':
		if eq.b!=0 {
			return eq.a / eq.b, nil
		} else{
			return 0, fmt.Errorf(DivisionByZeroError)
		}
	default: return 0, fmt.Errorf(IncorrectEquationError)
	}
}
