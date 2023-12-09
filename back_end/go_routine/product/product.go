package product

//This package realizes the class of Product, which is used to store the information of each detail of the product
//And it's included by outlet.go and central_hub.go

type Product struct {
	productID         string
	productName       string
	quantity          int
	replenishmentRate int //the number of products that can be replenished per day
	max_stock         int //the maximum number of products that can be stored
}

// Initialize with Values
func NewProduct(productID string, productName string, quantity int, replenishmentRate int, max_stock int) *Product {
	instance := &Product{
		productID:         productID,
		productName:       productName,
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

func (p Product) GetProductName() string {
	return p.productName
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

func (p *Product) SetProductName(productName string) {
	p.productName = productName
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
