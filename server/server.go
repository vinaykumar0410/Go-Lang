package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a user struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define a slice of users
var users []User
var nextID = 1

func main() {
	// Initialize with some sample users
	users = append(users, User{ID: nextID, Name: "Alice", Age: 30})
	nextID++
	users = append(users, User{ID: nextID, Name: "Bob", Age: 25})
	nextID++

	// Set up routes
	http.HandleFunc("/users", handleUsers) // All users
	http.HandleFunc("/users/", handleUser) // Specific user

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil)) 
	fmt.Println("Server is running on http://localhost:8080")
}

// Define handlers for each route 
func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Handle different HTTP methods
	switch r.Method {
	case http.MethodGet:
		// List all users
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		// Create a new user
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		user.ID = nextID
		nextID++
		users = append(users, user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}


func handleUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/users/"):])
	userIndex := -1
	for i, u := range users {
		if u.ID == id {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Get a specific user
		json.NewEncoder(w).Encode(users[userIndex])
	case http.MethodPut:
		// Update a user
		var updatedUser User
		json.NewDecoder(r.Body).Decode(&updatedUser)
		updatedUser.ID = id
		users[userIndex] = updatedUser
		json.NewEncoder(w).Encode(updatedUser)
	case http.MethodDelete:
		// Delete a user
		users = append(users[:userIndex], users[userIndex+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}