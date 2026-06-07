package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// --- DATA STRUCTURES ---

// Task 3 Struct
type Task3User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Mini Project 1 Struct
type MP1User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Mini Project 2 Struct
type MP2Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Mini Project 3 Structs
type MP3CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type MP3MessageResponse struct {
	Message string `json:"message"`
}

// Mini Project 4 Struct
type MP4Health struct {
	Status string `json:"status"`
}

// --- MAIN FUNCTION ---

func main() {
	// Task 1: Root Route
	http.HandleFunc("/", handleRootRoute)

	// Tasks 2, 3, 4 & Mini Projects 1, 3: Users Router
	http.HandleFunc("/users", handleUsersRouter)

	// Mini Project 2: Products Route
	http.HandleFunc("/products", handleProductsRoute)

	// Mini Project 4: Health Check Route
	http.HandleFunc("/health", handleHealthRoute)

	// Start the Server
	startServer()
}

// --- ROUTE HANDLERS & FUNCTIONS ---

// Task 1: Root Route Handler
func handleRootRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello Backend")
}

// Users Router (Orchestrates Task 2, 3, 4, MP1, and MP3)
func handleUsersRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Check for Task 4: '?page=' query parameter
		page := r.URL.Query().Get("page")

		if page != "" {
			// Executing Task 4 logging & Task 3 Response
			logTask4Query(page)
			respondWithTask3JSON(w)
		} else {
			// Executing Mini Project 1 (No page query provided)
			respondWithMP1Users(w)
		}

	case http.MethodPost:
		// Executing Mini Project 3
		respondWithMP3CreateUser(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Task 4: Log query parameter to server console
func logTask4Query(page string) {
	fmt.Printf("[Task 4] Requested Page: %s\n", page)
}

// Task 3: Send back single user payload
func respondWithTask3JSON(w http.ResponseWriter) {
	u := Task3User{Name: "Divyansh", Age: 21}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// Mini Project 1: Send back list of users
func respondWithMP1Users(w http.ResponseWriter) {
	users := []MP1User{
		{ID: 1, Name: "Divyansh"},
		{ID: 2, Name: "Rahul"},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Mini Project 3: Read post body and create user
func respondWithMP3CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser MP3CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(MP3MessageResponse{Message: "Invalid Payload"})
		return
	}

	// Print incoming post body data to console
	fmt.Printf("[Mini Project 3] Received Data -> Name: %s, Age: %d\n", newUser.Name, newUser.Age)

	w.WriteHeader(http.StatusCreated) // 201 Status
	json.NewEncoder(w).Encode(MP3MessageResponse{Message: "User Created"})
}

// Mini Project 2: Products Handler
func handleProductsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	products := []MP2Product{
		{ID: 1, Name: "Laptop"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Mini Project 4: Health Check Handler
func handleHealthRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MP4Health{Status: "ok"})
}

// Server Startup Utility
func startServer() {
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
