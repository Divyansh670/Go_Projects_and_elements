package main

import (
	"fmt"
	"net/http"
)

// Ek single string variable ko hi apna dummy database maan lete hain
var data = "Divyansh"

func handle(w http.ResponseWriter, r *http.Request) {
	// URL se 'name' parameter uthane ke liye (e.g., ?name=Rahul)
	name := r.URL.Query().Get("name")

	switch r.Method {
	case "GET":
		// 1. GET: Jo bhi data me save hai, screen par dikhao
		fmt.Fprintln(w, "Current Data:", data)

	case "POST":
		// 2. POST: Naya naam set karo
		data = name
		fmt.Fprintln(w, "Data Created/Set to:", data)

	case "PUT":
		// 3. PUT: Purane naam ko badal kar naya naam rakho
		data = name
		fmt.Fprintln(w, "Data Updated to:", data)

	case "DELETE":
		// 4. DELETE: Data ko khali kar do
		data = ""
		fmt.Fprintln(w, "Data Deleted!")
	}
}

func main() {
	http.HandleFunc("/users", handle)
	fmt.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}
