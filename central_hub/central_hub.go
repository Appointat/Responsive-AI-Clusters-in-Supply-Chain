package centralhub

import (
	"fmt"
	"sync"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

type Response struct {
	Replenishments map[string]int // Key is the name of the product, value is the number of products to replenish
	Error          error
}

type CentralHub struct {
	hubID      string
	location   string
	resources  map[string]*product.Product
	accessLock sync.Mutex
}

var (
	instance *CentralHub
	once     sync.Once
)

// Singleton Pattern
func GetHubInstanceDefault() *CentralHub {
	once.Do(func() {
		instance = &CentralHub{}
	})
	return instance
}

// Initialize with Values
func GetHubInstance(hubID string, location string) *CentralHub {
	once.Do(func() {
		instance = &CentralHub{
			hubID:     hubID,
			location:  location,
			resources: make(map[string]*product.Product),
		}
	})
	return instance
}

// Getters
func (h *CentralHub) GetHubID() string {
	return h.hubID
}

func (h *CentralHub) GetLocation() string {
	return h.location
}

func (h *CentralHub) GetNumberOfProducts(productName string) int {
	if product, ok := h.resources[productName]; ok {
		return product.GetNumber()
	}
	return -1 // Return -1 if the product is not found
}

// HandleEventNotification responds to an event by determining necessary product replenishments
func (h *CentralHub) HandleEventNotification(event string, shopInventory map[string]*product.Product) *Response {
	h.accessLock.Lock()
	defer h.accessLock.Unlock()

	replenishments := make(map[string]int)

	// Calculate the number of products that need to be replenished
	for name, prod := range shopInventory {
		quantityNeeded := prod.GetReplenishmentRate()
		if quantityNeeded > h.resources[name].GetNumber() {
			replenishments[name] = quantityNeeded
			h.resources[name].SetNumber(h.resources[name].GetNumber() - quantityNeeded)
		} else {
			replenishments[name] = 0 // No replenishment needed if sufficient stock is available
		}
	}

	if event != "" {
		fmt.Printf("Central Hub %s at %s received notification of event %s\n", h.hubID, h.location, event)
	}

	return &Response{
		Replenishments: replenishments,
		Error:          nil,
	}
}
