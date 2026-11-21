package main

import "fmt"

func main() {
	/*
		 * Your turn: Write a function that takes a []string of names and returns a map[string]int where each name maps to its length.
			* Then in main, print each name and its length.
			*
			* names := []string{"Melak", "Go", "Programming"}
			// should produce map: {"Melak": 5, "Go": 2, "Programming": 11}
	*/

	preNames := []string{"Melak", "Go", "Programming"}
	postNames := nameAndLengthPrinter(preNames)

	fmt.Println(postNames)
}

func nameAndLengthPrinter(names []string) map[string]int {
	mapNames := map[string]int{}

	for _, name := range names {
		mapNames[name] = len(name)
	}

	return mapNames
}
