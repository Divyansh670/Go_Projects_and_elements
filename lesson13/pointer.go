package main

import "fmt"

type User struct {
	name string
	age  int
}

func updateAge(user *User) {
	user.age = 22
}
func main() {
	var user User
	fmt.Println("Enter your name and age:")
	fmt.Scan(&user.name)
	fmt.Scan(&user.age)
	updateAge(&user)
	fmt.Println("Updated informations are:", user.name, user.age)
}
