package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ( NegativeInRomanError = "[!] В римской системе нет отрицательных чисел.")


func main() {
    fmt.Println("Calculator Application for kata academy")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input);
        eq, ansType, err := parseInput(input)
        if err != nil {
            fmt.Println(fmt.Sprint(err))
            continue
        }
        calculated, err := calculate(eq)
        if err != nil {
            fmt.Println(fmt.Sprint(err))
            continue
        }

        if calculated < 0 && ansType == Roman {
            fmt.Println(NegativeInRomanError)
            continue
        }
        
        ans, err := intToNumberType(ansType, calculated)

        if err != nil{
            fmt.Println(fmt.Sprint(err))
            continue
        }


        fmt.Println(ans)
	}
}
