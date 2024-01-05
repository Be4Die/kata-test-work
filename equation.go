package main

const (
	IncorrectEquationError = "[!] Строка не является математической операцией"
)

type Equation struct {
	a  int
	b  int
	op byte
}
