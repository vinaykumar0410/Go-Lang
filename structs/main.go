package main

import "fmt"

func main() {
	user := User{"Vinay", 23, "Kadapa"}
	fmt.Println(user)
	fmt.Printf("user is %+v \n", user)
	fmt.Printf("Name is %v \nAge is %v \nCity is %v", user.Name, user.Age, user.City)
	
}

type User struct {
	Name string
	Age  int
	City string
}
