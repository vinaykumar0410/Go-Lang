package main

import "fmt"

func main() {

	var friends [4]string

	fmt.Println("Friends are :", friends) // Friends are : [   ]

	friends[0] = "Vinay"
	friends[1] = "Hemanth"
	friends[3] = "Krishna"

	fmt.Println("Friends are :", friends) // [Vinay Hemanth  Krishna]

	fmt.Println(friends[2]) // ""

	fmt.Println("Length of friends array is :", len(friends)) // Length of friends array is : 4

	closeFriends := friends

	fmt.Println("Close friends are :", closeFriends) // Close friends are : [Vinay Hemanth  Krishna]

	if closeFriends == friends {
		fmt.Println("Arrays are same") // Arrays are same
	}

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are :", numbers) // Numbers are : [1 2 3 4 5]
}
