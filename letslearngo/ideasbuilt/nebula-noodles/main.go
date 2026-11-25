package main

import "fmt"

type Recipe struct {
	Name        string
	Ingredients map[string]int
	Price       int
	Tags        []string
}

type Customer struct {
	Name       string
	Preference string
}

type Ingredient struct {
}

type Game struct {
	Day        int
	Credits    int
	Reputation int
	Stock      map[string]int
	Menu       []Recipe
	Customers  []Customer
	Orders     []Order
}

type Order struct {
	Customer Customer
	Recipe   Recipe
}

func main() {

	game := Game{
		Day:        1,
		Credits:    0,
		Reputation: 0,
		Stock:      map[string]int{"noodles": 10, "space-sauce": 5, "stardust": 3, "protein": 8, "sweet-syrup": 4},
		Menu: []Recipe{
			{Name: "Menu1",
				Ingredients: map[string]int{"noodles": 2, "space-sauce": 1, "stardust": 1},
				Price:       10, Tags: []string{"spicy"}},
			{Name: "Menu2", Ingredients: map[string]int{"protein": 3, "sweet-syrup": 2}, Price: 15, Tags: []string{"sweet"}},
			{Name: "Menu3", Ingredients: map[string]int{"noodles": 3, "protein": 1}, Price: 25, Tags: []string{"expensive"}},
		},
		Customers: []Customer{
			{Name: "Melak", Preference: "spicy"},
			{Name: "Sisay", Preference: "sweet"},
			{Name: "Mekonen", Preference: "expensive"},
		},
	}

	fmt.Printf("Day %d | Credits: %d | Reputation: %d\n\n", game.Day, game.Credits, game.Reputation)
}
