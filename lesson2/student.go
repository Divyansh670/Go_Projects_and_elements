package main

import "fmt"

func main() {
	// First time introducing these variables -> Use :=
	name := "Divyansh"
	city := "Basti"
	const College = "KIET Group Of Institutions"
	age := 22

	fmt.Println(name)
	fmt.Println(city)
	fmt.Println(College)
	fmt.Println(age)

	// Updating existing variables -> Use =
	name = "Divyansh Srivastav"
	age = 23

	fmt.Println(name)
	fmt.Println(city)
	fmt.Println(College)
	fmt.Println(age)
}
