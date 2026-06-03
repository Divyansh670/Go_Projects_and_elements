package main

import "fmt"

func login(user map[string]string) bool {
	fmt.Println("Enter your name and password:")
	var name, password string
	fmt.Scan(&name)
	fmt.Scan(&password)
	c, exists := user[name]
	if exists && password == c {
		return true
	}

	return false
}

func deleteEmployee(salary map[string]int) {
	delete(salary, "Employee1")
}

func update(salary map[string]int) {
	salary["Employee2"] = 40000
	salary["Employee1"] = 30000
}

func main() {
	user := map[string]string{
		"admin": "12345",
		"user":  "divy12345",
	}
	log := login(user)
	if !log {
		fmt.Println("Login failed!")
		return
	}
	fmt.Println("Login successful!")
	salary := map[string]int{}
	salary["Employee1"] = 50000
	salary["Employee2"] = 50000
	salary["Employee3"] = 50000
	deleteEmployee(salary)
	update(salary)
	fmt.Println("Total number of Employees are Or length of the salary map is:", len(salary))
	fmt.Println("Employee and their salaries are:")
	for key, value := range salary {
		fmt.Println(key, ":", value)
	}
}
