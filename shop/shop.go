package shop

import (
	"sync"
	"time"
)

type Shop struct {
	shopID    string
	location  string
	inventory map[string]int
}

var (
	currentDate time.Time
	dateMutex   sync.Mutex
	once        sync.Once
)

// This function will be called automatically when the package is imported
func init() {
	once.Do(func() {
		//Initialize the currentData with the current time, such as today is...11/13/2023
		currentDate = time.Now()
		ticker := time.NewTicker(time.Second * 1)
		go func() {
			for {
				select {
				case <-ticker.C:
					dateMutex.Lock()
					//the currentDate
					currentDate = currentDate.AddDate(0, 0, 1)
					dateMutex.Unlock()
				}
			}
		}()
	})
}

func NewShop(shopID string, location string) *Shop {
	instance := &Shop{
		shopID:    shopID,
		location:  location,
		inventory: make(map[string]int),
	}

	return instance
}

func (s Shop) GetShopID() string {
	return s.shopID
}

func (s Shop) GetLocation() string {
	return s.location
}

func (s Shop) GetNumberOfProducts(NameofProduct string) int {
	if quantity, ok := s.inventory[NameofProduct]; ok {
		return quantity
	}

	//return -1 if the product is not found
	return -1
}

func GetCurrentDate() time.Time {
	dateMutex.Lock()
	defer dateMutex.Unlock()
	return currentDate
}
