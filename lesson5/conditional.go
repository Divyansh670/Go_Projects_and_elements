package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Yes, Go in")

		Hpassword := "Divy@1805"
		var password string
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		if Hpassword == password {
			fmt.Println("User")
		} else { 
			fmt.Println("Not a User")
		}

	} else { 
		fmt.Println("No, Get out")
	}
}
