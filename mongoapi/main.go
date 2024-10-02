package main

import (
	"fmt"
	"log"
	router "mongoapi/router"
	"net/http"
)

func main() {

	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	
	fmt.Println("Listening on port 8080")
}
