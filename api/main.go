package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func (user *User) IsEmpty() bool {
	// return user.Id == "" && user.Name == ""
	return user.Name == ""
}

var users []User

func main() {

	// init router
	r := mux.NewRouter()

	// seeding
	users = append(users, User{Id: "1",Name: "Vinay", Age: 23, Address: &Address{Street: "Fake st",City: "StandOut"}})
	users = append(users, User{Id: "2",Name: "Sai", Age: 22, Address: &Address{Street: "Real st",City: "StandIn"}})

	// routing
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/adduser", AddUser).Methods("POST")
	r.HandleFunc("/updateuser/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/deleteuser/{id}", DeleteUser).Methods("DELETE")


	// listen & serve 
	log.Fatal(http.ListenAndServe(":8080",r))


}

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Home Page"))

}


// get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

// get specific user
func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range users {
		if user.Id == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode("User Not Found With Specific Id")
}

// add user
func AddUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send Some Data")
	}

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	if user.IsEmpty() {
		json.NewEncoder(w).Encode("No Data Inside JSON")
		return 
	}

	user.Id = strconv.Itoa(rand.Intn(1000000)) 

	users = append(users, user)

	json.NewEncoder(w).Encode(user)

}	

// update user

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {
		if user.Id == params["id"] {

			var updatedUser User
			json.NewDecoder(r.Body).Decode(&updatedUser)

			if updatedUser.Name != "" {
				user.Name = updatedUser.Name
			}
			if updatedUser.Age != 0 {
				user.Age = updatedUser.Age
			}
			if updatedUser.Address != nil {
				user.Address = updatedUser.Address
			}

			users[index] = user

			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode("User Not Found With Specific Id")
}

// delete user

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index, user := range users {

		if user.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode("User Deleted")
			break
		}
	}
}