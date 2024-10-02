package main

import "fmt"

func main() {

	// creating developer object
	developer := Developer{"Vinay", true, 23, "Vinny.go"}

	fmt.Printf("Developer Details : %+v\n", developer)
	// Developer Details : {Name:Vinay TeamLead:true Age:23 Organization:Vinny.go}

	developer.getName()
	// Vinay

	developer.changeOrganization("Vinny.in")
	// Vinny.in

	fmt.Printf("Developer Details : %+v\n", developer)
	// Developer Details : {Name:Vinay TeamLead:true Age:23 Organization:Vinny.go}

}

// Initializing struct
type Developer struct {
	Name         string
	TeamLead     bool
	Age          int
	Organization string
}


// method to get name of developer

func (developer Developer) getName() {

	fmt.Println(developer.Name)

}

// method to modify the name of the organization
// here a copy of developer object is passed

func (developer Developer) changeOrganization(newOrganization string) {

	developer.Organization = newOrganization // not modifying original value

	fmt.Println(developer.Organization)
}