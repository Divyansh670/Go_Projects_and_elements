package main

import "fmt"

func main() {
	var sum int
	var n int
	fmt.Println("Enter Your number")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("The sum of all the numbers till your number is:", sum)
}
