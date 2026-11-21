package main

import "fmt"

func main() {
	/*
		 *
			* words := []string{"go", "is", "fun", "go", "is", "fast", "go", "rocks"}
	*/

	words := []string{"go", "is", "fun", "go", "is", "fast", "go", "rocks"}

	mappedWords := countWords(words)

	printCounts(mappedWords)
}

func countWords(words []string) map[string]int {

	var newWords = map[string]int{}

	for _, word := range words {
		_, ok := newWords[word]
		if ok {
			newWords[word] += 1
		} else {
			newWords[word] = 1
		}
	}

	return newWords
}

func printCounts(counts map[string]int) {
	for word, count := range counts {
		if count > 1 {
			fmt.Println(word, ":", count, "(repeated)")
		} else {
			fmt.Println(word, ":", count)
		}
	}
}
