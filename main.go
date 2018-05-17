package main

import (
	"fmt"
	"practice/go-calculator/calculator"
	"strconv"
)

func calculate(x int, y int, operation string) int {
	switch operation {
	case "ADD", "add":
		return calculator.Add(x, y)
	case "SUB", "sub":
		return calculator.Subtract(x, y)
	case "MUL", "mul":
		return calculator.Multiply(x, y)
	case "DIV", "div":
		res, _ := calculator.Divide(x, y)
		return res
	}
	fmt.Println("Please enter a valid operation. The operation will anyway return -1 in this case, which might not be the desired output")
	return -1
}

func main() {
	var s1, s2, s3 string

	for {
		fmt.Println("Enter your command (eg. 1 2 ADD, 2 3 SUB, 1 3 MUL, 6 2 DIV): ")
		fmt.Scanln(&s1, &s2, &s3)

		num1, err := strconv.Atoi(s1)
		if err != nil {
			fmt.Println("error detected")
		}

		num2, err := strconv.Atoi(s2)
		if err != nil {
			fmt.Println("error detected")
		}

		fmt.Println(calculate(num1, num2, s3))
	}
}
