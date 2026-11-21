package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	Pages     int
	Available bool
}

func (b Book) Summary() string {
	return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

func (b *Book) Borrow() bool {
	if b.Available {
		b.Available = false
		return true
	}

	return false

}

func (b *Book) Return() {
	b.Available = true
}
func main() {

	// The Library
	books := []Book{
		{
			Title:     "The Pragmatic Programmer",
			Author:    "David Thomas and Andrew Hunt",
			Pages:     352,
			Available: true,
		},
		{
			Title:     "Clean Code",
			Author:    "Robert C. Martin",
			Pages:     464,
			Available: true,
		},
		{
			Title:     "Designing Data-Intensive Applications",
			Author:    "Martin Kleppmann",
			Pages:     616,
			Available: false,
		},
		{
			Title:     "The Go Programming Language",
			Author:    "Alan A. A. Donovan and Brian W. Kernighan",
			Pages:     380,
			Available: true,
		},
	}

	// borroewdBy := map[string]string{} // title -> borrower name

	printLibrary(books)
	fmt.Println(findBook(books, "The Pragmatic Programmer"))

}

func findBook(books []Book, title string) *Book {
	for _, book := range books {
		if book.Title == title {
			return &book
		}
	}
	return nil
}

func borrowBook() {

}

func returnBook() {

}

func printLibrary(books []Book) {
	for _, book := range books {
		fmt.Println(book.Summary())
	}
}
