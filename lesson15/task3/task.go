package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by Zero.")
	}
	num := a / b
	return num, nil
}

func main() {
	num, err := divide(10, 0)
	fmt.Println("Error:", err)
	fmt.Println("Result:", num)
}
