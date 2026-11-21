package main

import "fmt"

func main() {
	/*
	 * Your turn: Write a function called greet that takes a name and prints "Hello, <name>!". Then call it from main with your name.
	 */

	name := "Melak"
	greet(name)
}

func greet(name string) {
	fmt.Printf("Hello, %s \n", name)
}
