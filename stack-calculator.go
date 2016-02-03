package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "1 2 + 4 * 5 + 3 -"
    operandStack := makeOperandStackFromString(str)
    numberStack := makeNumberStackFromString(str)
    fmt.Println(numberStack)
    fmt.Println(operandStack)

}

func makeNumberStackFromString(s string) []int{
    var numberStack []int
    for _, value := range s {
        if _, err := strconv.Atoi(string(value)); err == nil {                                                                                                                 
            numberStack = append(numberStack, int(value))
            fmt.Println(string(value))
        }
    }
    return numberStack
}

func makeOperandStackFromString(s string) []string {
	var operandStack []string
	for _, value := range s {
		if value == '+' || value == '-' || value == '*' || value == '/' {
			operandStack = append(operandStack, string(value))
			fmt.Println(string(value))
		} 
	}
	return operandStack
}