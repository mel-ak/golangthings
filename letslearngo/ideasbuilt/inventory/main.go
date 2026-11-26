package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {

	products := []Product{
		{Name: "Oil", Price: 2549.5, Quantity: 8},
		{Name: "Tissue", Price: 100, Quantity: 20},
		{Name: "NoteBook", Price: 75.5, Quantity: 5},
		{Name: "Pen", Price: 50, Quantity: 40},
		{Name: "Gum", Price: 25, Quantity: 255},
	}

	inventoryValue := 0.0
	countOfLowStocks := 0

	stock := map[string]float64{}

	fmt.Println("Inventory Report:")
	for _, product := range products {
		productValue := float64(product.Quantity) * product.Price
		stock[product.Name] = productValue
		inventoryValue += productValue
		status := "OK"
		if product.Quantity < 10 {
			status = "Low Stock"
			countOfLowStocks++
		}

		fmt.Printf("Name: %10s     | Price: %8.2f | Qty: %4d  | Total Value: %10.2f | Status: %s\n", product.Name, product.Price, product.Quantity, stock[product.Name], status)
	}

	fmt.Println("\n\n---")
	fmt.Printf("Total Inventory Value : %.2f\n", inventoryValue)
	fmt.Printf("Number of low stock items: %d\n", countOfLowStocks)

	if inventoryValue < 1000 {
		fmt.Println("Warning: Inventory value is unhealthy.")
	} else {
		fmt.Println("Inventory value is healthy.")
	}

}
