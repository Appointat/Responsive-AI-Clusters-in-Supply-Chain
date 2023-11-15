package manager

import (
	"fmt"
	"sync"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

type Response struct {
	Replenishments map[string]int //Key is the name of the product, value is the number of products that need to be replenished
	Error          error
}

type Manager struct {
	managerID  string
	location   string
	resources  map[string]*product.Product
	accessLock sync.Mutex
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
			resources: make(map[string]*product.Product),
		}
	})
	return instance
}

// Getters
func (m *Manager) GetManagerID() string {
	return m.managerID
}

func (m *Manager) GetLocation() string {
	//modify the value of location
	return m.location
}

func (m *Manager) GetNumebrOfProducts(NameofProduct string) int {
	if Product, ok := m.resources[NameofProduct]; ok {
		return Product.GetNumber()
	}
	//return -1 if the product is not found
	return -1
}

// With *, because we want to modify the value of the product
func (m *Manager) HandleEventNotification(event string, shopInventory map[string]*product.Product) *Response {
	m.accessLock.Lock()
	defer m.accessLock.Unlock()

	replenishments := make(map[string]int)

	//calculate the number of products that need to be replenished
	for name, prod := range shopInventory {
		quantityNeeded := prod.GetReplenishment_rate()
		if quantityNeeded > m.resources[name].GetNumber() {
			replenishments[name] = quantityNeeded
			m.resources[name].SetNumber(m.resources[name].GetNumber() - quantityNeeded)
			//if the stock in the manager's warehouse is enough, then replenish the shop
		} else {
			replenishments[name] = 0
		}
	}
	//deal with the event...if possible
	fmt.Println("Manager " + m.managerID + " at " + m.location + " received notification of event " + event)

	return &Response{
		Replenishments: replenishments,
		Error:          nil,
	}
}
