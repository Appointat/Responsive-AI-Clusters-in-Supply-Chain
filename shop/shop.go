package shop

import (
	"sync"
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

// Events string->time.Time
var HolidayEvents = map[string]time.Time{
	"New Year's Day":       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	"Valentine's Day":      time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
	"St. Patrick's Day":    time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
	"Easter":               time.Date(2023, 4, 4, 0, 0, 0, 0, time.UTC),
	"Independence Day":     time.Date(2023, 7, 4, 0, 0, 0, 0, time.UTC),
	"Halloween":            time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC),
	"Thanksgiving":         time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
	"Christmas Eve":        time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC),
	"Christmas Day":        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
	"New Year's Eve":       time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
	"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
}

type Shop struct {
	shopID    string
	location  string
	inventory map[string]product.Product
}

var (
	currentDate time.Time
	dateMutex   sync.Mutex
	once        sync.Once
)

// This function will be called automatically when the package is imported
func init() {
	//Initialize the HolidayEvents with the current year
	currentYear := time.Now().Year()
	for k, v := range HolidayEvents {
		HolidayEvents[k] = time.Date(currentYear, v.Month(), v.Day(), 0, 0, 0, 0, time.UTC)
	}
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

func GetCurrentDate() time.Time {
	dateMutex.Lock()
	defer dateMutex.Unlock()
	return currentDate
}

func NewShop(shopID string, location string) *Shop {
	instance := &Shop{
		shopID:    shopID,
		location:  location,
		inventory: make(map[string]product.Product),
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
	if Product, ok := s.inventory[NameofProduct]; ok {
		return Product.GetNumber()
	}

	//return -1 if the product is not found
	return -1
}

//TODO: Add a method to check all the stock of the products in the shop and trigger the replenishment function based on each product's replenishment rate
//Every day, the shop will loop through all the products in the shop and trigger the replenishment function
