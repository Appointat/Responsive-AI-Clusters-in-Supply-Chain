package manager

import (
	"fmt"
	"sync"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

type Manager struct {
	managerID string
	location  string
	resources map[string]product.Product
}

var (
	instance *Manager
	once     sync.Once
)

// Singleton Pattern
func GetManagerInstanceDefault() *Manager {
	once.Do(func() {
		instance = &Manager{}
	})
	return instance
}

// Initialize with Values
func GetManagerInstance(managerID string, location string) *Manager {
	once.Do(func() {
		instance = &Manager{
			managerID: managerID,
			location:  location,
			resources: make(map[string]product.Product),
		}
	})
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

// With *, because we want to modify the value of the product
func (m *Manager) HandleEventNotification(event string) {
	fmt.Println("Manager " + m.managerID + " at " + m.location + " received notification of event " + event)
}
