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
			Available: true,
		},
		{
			Title:     "The Go Programming Language",
			Author:    "Alan A. A. Donovan and Brian W. Kernighan",
			Pages:     380,
			Available: true,
		},
	}

	borrowedBy := map[string]string{} // title -> borrower name

	borrower1 := "Melak"
	borrower2 := "Bob"
	sampleBook := "Clean Code"

	fmt.Println("LIBRARY :\n")
	printLibrary(books, borrowedBy)

	fmt.Printf("\n\nBORROW : (%s - %s)\n", borrower1, sampleBook)
	err := borrowBook(books, borrowedBy, sampleBook, borrower1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully")
	}

	err = borrowBook(books, borrowedBy, sampleBook, borrower2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully")
	}

	fmt.Println("LIBRARY :\n")
	printLibrary(books, borrowedBy)

	err = returnBook(books, borrowedBy, sampleBook)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully")
	}

	err = returnBook(books, borrowedBy, sampleBook)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully")
	}

	fmt.Println("LIBRARY :\n")
	printLibrary(books, borrowedBy)
	// fmt.Println("\n\nSearching Book ...", sampleBook, "...")
	// fmt.Println(findBook(books, "The Pragmatic Programmer"))

}

func findBook(books []Book, title string) *Book {
	for i, book := range books {
		if book.Title == title {
			return &books[i]
		}
	}
	return nil
}

func borrowBook(books []Book, borrowedBy map[string]string, title, person string) error {

	book := findBook(books, title)

	if book == nil {
		return fmt.Errorf("book %q not found", title)
	}

	if book.Available == false {
		return fmt.Errorf("%q is already borrowed", title)
	}

	if !book.Borrow() {
		return fmt.Errorf("Error borrowing")
	}
	borrowedBy[title] = person

	return nil


}

func returnBook(books []Book, borrowedBy map[string]string, title string) error {

	book := findBook(books, title)
	if book == nil {
		return fmt.Errorf("book %q not found", title)
	}

	if borrowedBy[title] == "" {
		return fmt.Errorf("%q was not borrowed", title)
	}

	book.Return()

	delete(borrowedBy, title)

	return nil
}

func printLibrary(books []Book, borrowedBy map[string]string) {
	for _, book := range books {
		state := ""
		if book.Available {
			state = "available"
		} else {
			state = "borrowed by " + borrowedBy[book.Title]
		}
		fmt.Println(book.Summary(), "-", state)
	}
}
