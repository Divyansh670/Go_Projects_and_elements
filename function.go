package main

import "fmt"

func greet(name string) {
	fmt.Println("Welcome Mr.", name)
}
func square(num int) int {
	return num * num
}
func main() {
	for i := 1; i < 4; i++ {
		greet("Divy")
	}
	var n int
	fmt.Println("Enter the number that is to be squared:")
	fmt.Scan(&n)
	var sq int = square(n)
	fmt.Println("The square of the number is:", sq)
}
