package product

//This package realizes the class of Product, which is used to store the information of each detail of the product
//And it's included by shop.go and manager.go

type Product struct {
	productName        string
	Number             int
	replenishment_rate int //the number of products that can be replenished per day
	max_stock          int //the maximum number of products that can be stored
}

// Initialize with Values
func NewProduct(productName string, Number int, replenishment_rate int, max_stock int) *Product {
	instance := &Product{
		productName:        productName,
		Number:             Number,
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
	return p.Number
}

func (p Product) GetReplenishment_rate() int {
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
	p.Number = Number
}

func (p *Product) SetReplenishment_rate(replenishment_rate int) {
	p.replenishment_rate = replenishment_rate
}

func (p *Product) SetMax_stock(max_stock int) {
	p.max_stock = max_stock
}

//TODO: Add a method to inform the manager of the number of products that need to be replenished per day
