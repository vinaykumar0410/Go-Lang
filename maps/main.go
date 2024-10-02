package main

import "fmt"

func main() {

	var heights = make(map[string]int)

	heights["Vinay"] = 157
	heights["Hemanth"] = 162
	heights["Krishna"] = 148
	heights["Hari"] = 154

	fmt.Println(heights) // map[Hari:154 Hemanth:162 Krishna:148 Vinay:157]
	
	fmt.Println("Height of Hemanth",heights["Hemanth"]) // Height of Hemanth 162
	
	delete(heights, "Hari")
	
	fmt.Println(heights) // map[Hemanth:162 Krishna:148 Vinay:157]

	// iterate over map

	for key, value := range heights {
		fmt.Printf("Height of %v is %vcms \n", key, value)
	}

	/*
		Height of Vinay is 157cms 
		Height of Hemanth is 162cms 
		Height of Krishna is 148cms 

	*/ 

	fmt.Println("Map length is",len(heights)) // 3

	// can delete all heights 
	for k := range heights {
		delete(heights, k)
	}
	
	fmt.Println("Map length is",len(heights)) // 0

}
