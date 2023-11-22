package product

//This package realizes the class of Product, which is used to store the information of each detail of the product
//And it's included by outlet.go and central_hub.go

type Product struct {
	productName        string
	number             int
	replenishment_rate int //the number of products that can be replenished per day
	max_stock          int //the maximum number of products that can be stored
}

// Initialize with Values
func NewProduct(productName string, number int, replenishment_rate int, max_stock int) *Product {
	instance := &Product{
		productName:        productName,
		number:             number,
		replenishment_rate: replenishment_rate,
		max_stock:          max_stock,
	}
	return instance
}

// Getters
func (p Product) GetProductName() string {
	return p.productName
}

func (p Product) GetNumber() int {
	return p.number
}

func (p Product) GetReplenishmentRate() int {
	return p.replenishment_rate
}

func (p Product) GetMax_stock() int {
	return p.max_stock
}

// Setters
func (p *Product) SetProductName(productName string) {
	p.productName = productName
}

func (p *Product) SetNumber(Number int) {
	p.number = Number
}

func (p *Product) SetReplenishmentRate(replenishment_rate int) {
	p.replenishment_rate = replenishment_rate
}

func (p *Product) SetMaxStock(max_stock int) {
	p.max_stock = max_stock
}

//TODO: Add a method to inform the central hub of the number of products that need to be replenished per day
