// Assignment 1
//author - Meet Mehta

package main

import "fmt"

func main() {

	//fmt.Scan()
	//fmt.Scanln()
	//fmt.Scanf()
	//can use any of the above 3 methods to take user input

	var name string //variable for storing name string

	fmt.Print("what is your name ? ")
	fmt.Scan(&name)

	fmt.Printf("hello %s", name) //%s is for string
}
