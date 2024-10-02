package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// var myurl = "https://vinay-flutter.netlify.app/"
func main() {
	/*
	response, err := http.Get(myurl) 
	handleError(err)

	databytes, err := io.ReadAll(response.Body) 

	defer response.Body.Close() 
	handleError(err)

	content := string(databytes)

	fmt.Println(content) 
	*/

	GetRequest("http://localhost:5000") // Content is <center> <h1> Express as Backend Server for Go Lang </h1> </center>
	GetRequest("http://localhost:5000/get") // Content is {"message":"Hello from backend server This is Get Request"}
	PostRequest("http://localhost:5000/post") // Content is {"message":"This is Post Request","statuscode":"200","success":"true"}
	PostFormRequest("http://localhost:5000/postform") // Content is {"name":"vinay","age":"23","city":"kadapa"}
}

func handleError(err error){
	
	if err != nil {
		panic(err)
	}

}

func GetRequest(endpoint string){

	response, err := http.Get(endpoint)
	handleError(err)
	
	defer response.Body.Close()
	
	content , err := io.ReadAll(response.Body);
	handleError(err)

	// fmt.Println("Content is", string(content)) // short way

	var builder strings.Builder
	builder.Write(content)

	fmt.Println("Content is", builder.String())

}

func PostRequest(endpoint string) {
	
	requestBody := strings.NewReader(`
	{
		"message" : "This is Post Request",
		"statuscode" : "200",
		"success" : "true"
	}
	`)

	response, err := http.Post(endpoint, "application/json", requestBody)
	handleError(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	handleError(err)

	// fmt.Println("Content is", string(content))

	var builder strings.Builder
	builder.Write(content)

	fmt.Println("Content is", builder.String())

}

func PostFormRequest(endpoint string){

	data := url.Values{}

	data.Add("name", "vinay")
	data.Add("age", "23")
	data.Add("city", "kadapa")

	response, err := http.PostForm(endpoint, data)
	handleError(err)

	content, err := io.ReadAll(response.Body)
	handleError(err)

	// fmt.Println("Content is", string(content))

	var builder strings.Builder
	builder.Write(content)

	fmt.Println("Content is", builder.String())
}