package product

import "time"

//This package realizes the class of Product, which is used to store the information of each detail of the product
//And it's included by outlet.go and central_hub.go

type Product struct {
	productID         string
	quantity          int
	replenishmentRate int //the number of products that can be replenished per day
	max_stock         int //the maximum number of products that can be stored
}

// HolidayEvents maps event names to their dates.
var HolidayMaps = []map[string]time.Time{
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	// Copy the same map for the sake of this example
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	// Repeat the same map 3 more times to simulate different maps
	// In real usage, these would likely be different
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
}
var GlobalProducts []Product

func InstanceProducts() {
	GlobalProducts = []Product{
		{productID: "Olive Oil", quantity: 100, replenishmentRate: 30, max_stock: 500},
		{productID: "Baguette", quantity: 200, replenishmentRate: 50, max_stock: 300},
		{productID: "Manchego Cheese", quantity: 150, replenishmentRate: 40, max_stock: 400},
		{productID: "Black Tea", quantity: 80, replenishmentRate: 20, max_stock: 250},
	}
	// Add any additional logic if required
}

// Initialize with Values
func NewProduct(productID string, quantity int, replenishmentRate int, max_stock int) *Product {
	instance := &Product{
		productID:         productID,
		quantity:          quantity,
		replenishmentRate: replenishmentRate,
		max_stock:         max_stock,
	}
	return instance
}

// Getters
func (p Product) GetProductID() string {
	return p.productID
}

func (p Product) GetNumber() int {
	return p.quantity
}

func (p Product) GetReplenishmentRate() int {
	return p.replenishmentRate
}

func (p Product) GetMax_stock() int {
	return p.max_stock
}

// Setters
func (p *Product) SetProductID(productID string) {
	p.productID = productID
}

func (p *Product) SetNumber(Quantity int) {
	p.quantity = Quantity
}

func (p *Product) SetReplenishmentRate(replenishment_rate int) {
	p.replenishmentRate = replenishment_rate
}

func (p *Product) SetMaxStock(max_stock int) {
	p.max_stock = max_stock
}

//TODO: Add a method to inform the central hub of the number of products that need to be replenished per day
