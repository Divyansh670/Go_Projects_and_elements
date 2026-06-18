package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Task 1 & 2: User struct with lower-case JSON tags
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("=== RUNNING ALL TASKS ===")

	task1And2()
	task3()
	task4()
}

// ==========================================
// TASK 1 & 2: Marshal Object with JSON Tags
// ==========================================
func task1And2() {
	fmt.Println("\n--- Task 1 & 2: Marshal to JSON ---")

	u1 := User{
		Name: "Divyansh",
		Age:  21,
	}

	// MarshalIndent se JSON pretty-print (format) ho kar aata hai
	jsonData, err := json.MarshalIndent(u1, "", "  ")
	if err != nil {
		log.Fatalf("Task 1&2 Error: %s", err)
	}

	fmt.Println(string(jsonData))
}

// ==========================================
// TASK 3: Unmarshal JSON String into Struct
// ==========================================
func task3() {
	fmt.Println("\n--- Task 3: Unmarshal from JSON ---")

	// Raw JSON string input
	jsonInput := `{"name": "Divyansh", "age": 21}`

	// Khali struct jisme data store hoga
	var u2 User

	// string ko []byte me badal kar pointer (&u2) ke saath pass kiya
	err := json.Unmarshal([]byte(jsonInput), &u2)
	if err != nil {
		log.Fatalf("Task 3 Error: %s", err)
	}

	// Alag-alag fields ko print kiya
	fmt.Printf("Extracted Fields -> Name: %s, Age: %d\n", u2.Name, u2.Age)
}

// ==========================================
// TASK 4: Convert Slice of Structs to JSON
// ==========================================
func task4() {
	fmt.Println("\n--- Task 4: Slice of Users to JSON ---")

	// []User yaani slice of User structs banaya
	users := []User{
		{Name: "Divyansh", Age: 21},
		{Name: "Anjali", Age: 23},
	}

	// Slice ko JSON Array me badla
	sliceData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalf("Task 4 Error: %s", err)
	}

	fmt.Println(string(sliceData))
}
