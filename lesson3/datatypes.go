package main

import "fmt"

func main() {
	name := "divy"
	fmt.Printf("Value: %s | Type: %T\n", name, name)

	age := 23
	fmt.Printf("Value: %d | Type: %T\n", age, age)

	gpa := 8.5
	fmt.Printf("Value: %f | Type: %T\n", gpa, gpa)
	fmt.Printf("Value: %.1f | Type: %T\n", gpa, gpa)

	active := true
	fmt.Printf("Value: %t | Type: %T\n", active, active)

	grade := 'A'
	fmt.Printf("Value: %c | Type: %T\n", grade, grade)
}
