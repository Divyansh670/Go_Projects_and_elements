package main

import "fmt"

func op(a, b int) (int, int) {
	return a + b, a * b
}

func getUser() (string, int) {
	name := "Divyansh"
	age := 23
	return name, age
}

func main() {
	var n1 int
	var n2 int
	fmt.Println("Enter Your both numbers for operations:")
	fmt.Scan(&n1)
	fmt.Scan(&n2)

	name, _ := getUser()
	fmt.Println("My name is:", name)
	sum, product := op(n1, n2)
	fmt.Println("The sum of both numbers are:", sum, " and the product of both numbers are:", product)
}
