package main

import "fmt"

type User struct {
	Name  string
	Age   int
	ID    int
	Email string
}

func withoutUserInput() {
	fmt.Println("--- Running Without User Input ---")
	var users []User

	users = append(users, User{Name: "Alice", Age: 25, ID: 101, Email: "alice@example.com"})
	users = append(users, User{Name: "Bob", Age: 30, ID: 102, Email: "bob@example.com"})
	users = append(users, User{Name: "Charlie", Age: 22, ID: 103, Email: "charlie@example.com"})

	users[1].ID = 999

	fmt.Println("All users in the slice:")
	for index, user := range users {
		fmt.Printf("Index %d -> ID: %d, Name: %s, Age: %d, Email: %s\n", index, user.ID, user.Name, user.Age, user.Email)
	}
}

func withUserInput() {
	fmt.Println("\n--- Running With User Input ---")
	var users []User
	var totalUsers int

	fmt.Print("Aapko kitne users ka data enter karna hai?: ")
	fmt.Scan(&totalUsers)

	for i := 0; i < totalUsers; i++ {
		var u User
		fmt.Printf("\nEnter details for User %d (Name, Age, ID, Email):\n", i+1)
		fmt.Scan(&u.Name, &u.Age, &u.ID, &u.Email)
		users = append(users, u)
	}

	if len(users) >= 2 {
		users[1].ID = 999
	}

	fmt.Println("\nAll users in the slice:")
	for index, user := range users {
		fmt.Printf("Index %d -> ID: %d, Name: %s, Age: %d, Email: %s\n", index, user.ID, user.Name, user.Age, user.Email)
	}
}

func main() {
	withoutUserInput()
	withUserInput()
}
