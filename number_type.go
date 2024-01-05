package main

const (
	UnknownTypeError = "[!] Неизвестный тип числа"
)

type NumberType int64

const (
	Arabic NumberType = iota
	Roman
)