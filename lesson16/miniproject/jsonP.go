package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// ==========================================
// STRUCT DEFINITIONS (Saare projects ke liye)
// ==========================================

// Mini Project 1 Struct
type Student struct {
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

// Mini Project 2 Struct
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Mini Project 3 Struct
type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Mini Project 4 Struct
type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// ==========================================
// MAIN FUNCTION (Sabhi projects ko chalane ke liye)
// ==========================================
func main() {
	fmt.Println("======== RUNNING MINI PROJECTS ========")

	miniProject1()
	miniProject2()
	miniProject3()
	miniProject4()
}

// ==========================================
// Mini Project 1: Student JSON (Slice of Structs)
// ==========================================
func miniProject1() {
	fmt.Println("\n--- Mini Project 1: Student JSON ---")

	// 3 Students ka slice banaya
	students := []Student{
		{Name: "Divyansh", Marks: 95},
		{Name: "Rahul", Marks: 88},
		{Name: "Amit", Marks: 76},
	}

	// JSON me marshal kiya
	jsonData, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		log.Fatalf("Mini Project 1 Error: %s", err)
	}

	fmt.Println(string(jsonData))
}

// ==========================================
// Mini Project 2: Product Catalog JSON
// ==========================================
func miniProject2() {
	fmt.Println("\n--- Mini Project 2: Product Catalog ---")

	// 3 Products ka slice banaya
	products := []Product{
		{ID: 101, Name: "Laptop", Price: 55000},
		{ID: 102, Name: "Smartphone", Price: 25000},
		{ID: 103, Name: "Headphones", Price: 3000},
	}

	// Slice ko JSON array me badla
	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatalf("Mini Project 2 Error: %s", err)
	}

	fmt.Println(string(jsonData))
}

// ==========================================
// Mini Project 3: User Registration Request (Unmarshal)
// ==========================================
func miniProject3() {
	fmt.Println("\n--- Mini Project 3: User Registration ---")

	// Input JSON string
	jsonInput := `{
	  "name": "Divyansh",
	  "email": "div@gmail.com"
	}`

	var req RegisterRequest

	// JSON string ko struct me unmarshal kiya (Pointer '&req' ke saath)
	err := json.Unmarshal([]byte(jsonInput), &req)
	if err != nil {
		log.Fatalf("Mini Project 3 Error: %s", err)
	}

	// Struct ke alag-alag fields ko print kiya
	fmt.Printf("Extracted Data -> Name: %s | Email: %s\n", req.Name, req.Email)
}

// ==========================================
// Mini Project 4: API Response Builder
// ==========================================
func miniProject4() {
	fmt.Println("\n--- Mini Project 4: API Response Builder ---")

	// Response object create kiya
	res := Response{
		Message: "User Created",
		Success: true,
	}

	// Object ko JSON me convert kiya
	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatalf("Mini Project 4 Error: %s", err)
	}

	fmt.Println(string(jsonData))
}
