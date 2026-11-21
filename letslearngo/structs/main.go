package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) Summary() string {
	return fmt.Sprintf("%s by %s (%d)", b.Title, b.Author, b.Pages)
}

// Value Receiver (just a copy)
func (p Person) Greet() string {
	return "Hi I'm " + p.Name
}

// Pointer Reciever (can change the original)
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {

	p1 := Person{Name: "Melak", Age: 45}

	fmt.Printf("My name is %s and age is %d\n", p1.Name, p1.Age)

	p2 := Person{"Bob", 25} // Not recommended

	fmt.Printf("My name is %s and age is %d\n", p2.Name, p2.Age)

	// Testing attached methods

	alice := Person{Name: "Alice", Age: 30}
	fmt.Println(alice.Greet())

	alice.HaveBirthday()
	fmt.Println(alice.Age)

	books := []Book{
		{Title: "Abc", Author: "melak", Pages: 10},
		{Title: "Abcd", Author: "melake", Pages: 11},
		{Title: "Abce", Author: "melakie", Pages: 12},
	}

	for _, book := range books {
		fmt.Println(book.Summary())
	}

}
