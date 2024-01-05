package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MaxRomanNumber = 3999
	MaxRomanNumberError = "[!] Невозможно перевести число больше 3999 в римскую систему исчисления"

	// В системе римских цифр отсутствует ноль, но ранее использовалось обозначение нуля как nulla (нет), nihil (ничто) и N (первая буква этих слов). (https://en.wikipedia.org/wiki/Roman_numerals#:~:text=numeral.%5B40%5D-,Zero,the%2020th%20century%20to%20designate%20quantities%20in%20pharmaceutical%20prescriptions.%5B44%5D,-Fractions)
	ZeroInRoman = "nulla"
	RomanToIntConvertError = "[!] Невозможно конвертировать данное число из римского в арабское"
)

func intToNumberType (numType NumberType, n int) (string, error) {
	switch numType {
	case Arabic: return strconv.Itoa(n), nil
	case Roman: 
		s, err := intToRoman(n)
		if err != nil{
			return "", err
		}
		return s, nil
	default:
		return "", fmt.Errorf(UnknownTypeError)
	}

}

func convertNumber(number string) (int, NumberType, error) {
	i, err := strconv.Atoi(number)
    if err != nil {
        fromRoman, err := romanToInt(number)
		if err != nil {
			return 0, 0,err
		}
		
		return fromRoman, Roman,nil
    }

	return i, Arabic, nil
}

func intToRoman(num int) (string, error) {
	// check limit
	if num > MaxRomanNumber {
		return "", fmt.Errorf(MaxRomanNumberError)
	}

	//check zero
	if (num == 0){
		return ZeroInRoman, nil
	}

	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var roman strings.Builder
	for num > 0{
        for i := range value{
            if num >= value[i]{
                roman.WriteString(symbol[i])
                num -= value[i]
                break
            }
        }
    }

	return roman.String(), nil
}

func romanToInt(roman string) (int, error) {
	if roman == ZeroInRoman{
		return 0, nil
	}

	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(roman) - 1; i >= 0; i-- {
		tmp, ok := m[roman[i]]
		if !ok{
			return 0, fmt.Errorf(RomanToIntConvertError)
		}
		sign := 1
		if tmp < last {
			sign = -1
		}
		res += sign * tmp
		last = tmp
	}

	if res > MaxRomanNumber {
		return 0, fmt.Errorf(MaxRomanNumberError)
	}
	return res, nil
}