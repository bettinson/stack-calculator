package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bettinson/stack-calculator/stack"
)

func main() {
	s := stack.New()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		operandString, _ := reader.ReadString('\n')
		operandString = operandString[:len(operandString)-1]
		if operandString == "+" || operandString == "-" || operandString == "*" || operandString == "/" {
			firstVal, firstErr := s.Pop()
			secondVal, secondErr := s.Pop()
			if firstErr != nil || secondErr != nil {
				log.Fatal(firstErr, secondErr)
			}
			var result int
			switch {
			case operandString == "+":
				result = firstVal + secondVal
			case operandString == "-":
				result = firstVal - secondVal
			case operandString == "*":
				result = firstVal * secondVal
			case operandString == "/":
				result = firstVal / secondVal
			default:
				log.Fatal()
			}
			fmt.Println(result)
			s.Push(result)
			continue
		}
		if operandString[0] >= '0' || operandString[0] <= '9' {
			operandInt, err := strconv.Atoi(operandString[:])
			if err != nil {
				log.Fatal(err)
			}
			s.Push(operandInt)
		}
	}
}
