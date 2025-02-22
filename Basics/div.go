package main

import (
	"fmt"
	"log"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var num1, num2 float64

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	result, err := divide(num1, num2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The result of %.2f divided by %.2f is %.2f\n", num1, num2, result)
}
