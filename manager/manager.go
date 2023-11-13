package manager

import (
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

type Manager struct {
	managerID string
	location  string
	resources map[string]product.Product
}

// Initialize with Values
func NewManager(managerID string, location string) *Manager {
	instance := &Manager{
		managerID: managerID,
		location:  location,
		resources: make(map[string]product.Product),
	}
	return instance
}

// Getters
func (m Manager) GetManagerID() string {
	return m.managerID
}

func (m Manager) GetLocation() string {
	//modify the value of location
	return m.location
}

func (m Manager) GetNumebrOfProducts(NameofProduct string) int {
	if Product, ok := m.resources[NameofProduct]; ok {
		return Product.GetNumber()
	}
	//return -1 if the product is not found
	return -1
}

//TODO: Add a method to accept the request of the replenishment from the shop
