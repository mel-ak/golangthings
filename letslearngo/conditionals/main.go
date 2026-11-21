package main

import "fmt"

func main() {
	/*
	 * Your turn: Write a program that prints the numbers 1 to 10, but only prints even numbers. Use a loop and an if statement together.
	 */

	var uptolimit uint = 10

	evenumbers(uptolimit)
}

func evenumbers(limit uint) {
	var i uint
	for i = 1; i <= limit; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		}
	}
	fmt.Println("Printing 1,000,000 USD money")
}
