package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"name"` 
	Age      int `json:"age"`
	Email    string `json:"email"`
	Password string `json:"-" ` // it will not be shown
	Skills   []string `json:"skills,omitempty"` // if skills is empty then it will not be shown
}

func main() {

	users := []user{
		{
			"vinay", 23, "vinay@go.dev", "abc123", []string{"Go", "JavaScript"},
		},
		{
			"sai", 22, "sai@go.dev", "xyz123", []string{"Python", "C"},
		},
		{
			"ram", 21, "ram@go.dev", "def123", nil,
		},
	}

	EndodeJson(users)

	DecodeJson()
}

func EndodeJson(users []user) {
	// fmt.Println("List of users :", users) // gives only values

	jsonFormat, err := json.Marshal(users) // gives json format in bytes
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(jsonFormat)) // gives json format with key values

	formatedJson, err := json.MarshalIndent(users,"","\t") // gives json format in bytes

	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(formatedJson)) // gives json format with key values and indentation
}

func DecodeJson(){

	jsonFromWeb := []byte(`
		{
                "name": "vinay",
                "age": 23,
                "email": "vinay@go.dev",
                "skills": [
                        "Go",
                        "JavaScript"
                ]
        }
	`)

	var vinay user;

	valid := json.Valid(jsonFromWeb) // returns true if json is valid

	if valid {
		json.Unmarshal(jsonFromWeb, &vinay) // unmarshalling the json
		fmt.Printf("%#v\n", vinay) 
	} else {
		fmt.Println("Invalid Json Format")
	}

	// convert json to map
	var jsonMap map[string]interface{}

	json.Unmarshal(jsonFromWeb, &jsonMap)

	fmt.Printf("%#v\n", jsonMap)

	for key, value := range jsonMap {
		fmt.Printf("Key is : %v & Value is : %v & Type is : %T\n", key, value, value)
	}
}