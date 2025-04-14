package inventory

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// Item represents a product in the warehouse
type Item struct {
	ID       string
	Name     string
	Quantity int
}

// InventoryService handles inventory operations
type InventoryService struct {
	items map[string]Item
}

// NewInventoryService creates a new inventory service
func NewInventoryService() *InventoryService {
	// Sample data
	return &InventoryService{
		items: map[string]Item{
			"P001": {ID: "P001", Name: "Laptop", Quantity: 15},
			"P002": {ID: "P002", Name: "Mouse", Quantity: 5},
		},
	}
}

// CheckStock verifies the stock
func (s *InventoryService) CheckStock(id string) (Item, error) {
	logger.Info("Checking stock", "product_id", id)
	item, exists := s.items[id]
	if !exists {
		logger.Error("Product not found", "product_id", id)
		return Item{}, errors.New("product not found")
	}
	if item.Quantity < 10 {
		logger.Warn("Low stock detected", "product_id", id, "quantity", item.Quantity)
		fmt.Printf("⚠️ Warning: Low stock for %s (Quantity: %d)\n", item.Name, item.Quantity)
	} else {
		logger.Info("Stock sufficient", "product_id", id, "quantity", item.Quantity)
	}
	return item, nil
}
