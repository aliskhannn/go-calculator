package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
    var a, b int
    var operator string

    fmt.Print("Enter the first number: ")
    _, err := fmt.Scan(&a)

    if err != nil {
        fmt.Println("Invalid input for the first number. Please enter an integer")
    }

    fmt.Print("Enter an operator (+, -, * or /): ")
    _, err = fmt.Scan(&operator)

    if err != nil || !isValidOperator(operator) {
        fmt.Println("Invalid operator. Please use one of +, -, * or /")
    }

    fmt.Print("Enter the second number: ")
    _, err = fmt.Scan(&b)

    if err != nil {
        fmt.Println("Invalid input for the second number. Please enter an integer")
    }

    switch operator {
    case "+":
        fmt.Printf("%v + %v = %v\n", a, b, add(a, b))
    case "-":
        fmt.Printf("%v - %v = %v\n", a, b, subtract(a, b))
    case "*":
        fmt.Printf("%v * %v = %v\n", a, b, multiply(a, b))
    case "/":
        result, err := divide(a, b)

        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("%v / %v = %v\n", a, b, result)
        }
    default:
        fmt.Println("Invalid operator")
    }
}

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("you can't divide by zero")
    }

    return a / b, nil
}

func isValidOperator(operator string) bool {
    validOperators := []string{"+", "-", "*", "/"}
    operator = strings.TrimSpace(operator)

    for _, validOprtr := range validOperators {
        if validOprtr == operator {
            return true
        }
    }

    return false
}
