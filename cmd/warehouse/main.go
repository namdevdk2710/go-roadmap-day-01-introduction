package main

import (
	"bufio"
	"fmt"
	"os"
	"warehouse-cli/internal/inventory"
)

func main() {
	fmt.Println("Warehouse CLI: Enter product ID to check stock (or 'exit' to quit)")

	service := inventory.NewInventoryService()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Product ID: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		item, err := service.CheckStock(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Product: %s, Quantity: %d\n", item.Name, item.Quantity)
	}
}
