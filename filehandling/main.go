package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Create("./hello.txt")
	handleErr(err)

	content := "Hello I am Learning File Handling using GO Language"

	length, err := file.WriteString(content)
	handleErr(err)

	fmt.Println("Length of content is", length)

	ReadFile("./hello.txt")

	file.Close()
}

func ReadFile(filepath string) {

	text, err := ioutil.ReadFile(filepath)
	handleErr(err)

	fmt.Println(string(text))
}

func handleErr(err error) {

	if err != nil {
		panic(err)
	}

}
