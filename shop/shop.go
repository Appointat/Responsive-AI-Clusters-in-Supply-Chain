package shop

import (
	"math/rand"
	"sync"
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/manager"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

// Events string->time.Time
var HolidayEvents = map[string]time.Time{
	"New Year's Day":       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	"Valentine's Day":      time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
	"St. Patrick's Day":    time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
	"Easter":               time.Date(2023, 4, 4, 0, 0, 0, 0, time.UTC),
	"Halloween":            time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC),
	"Thanksgiving":         time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
	"Christmas Eve":        time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC),
	"Christmas Day":        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
	"New Year's Eve":       time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
	"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
	"Test1":                time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
}

type Shop struct {
	shopID   string
	location string

	inventory map[string]*product.Product

	numberofEvents int      //the number of events that the shop is observing
	observedEvents []string //the events that the shop is observing

	notifyManager func(string, map[string]*product.Product) *manager.Response
}

var (
	currentDate time.Time
	dateMutex   sync.Mutex
	once        sync.Once
	wg          sync.WaitGroup
	ticker      *time.Ticker
	allShops    []*Shop
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
		ticker = time.NewTicker(time.Second * 1)
		go func() {
			for {
				select {
				case <-ticker.C:
					dateMutex.Lock()
					//the currentDate
					currentDate = currentDate.AddDate(0, 0, 1)
					dateMutex.Unlock()

					wg.Add(len(allShops))
					for _, shop := range allShops {
						go shop.CheckAndNotify(currentDate)

					}

					//Wait before all the shops are notified
					wg.Wait()
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

func NewShop(shopID string, location string, numberofEvents int) *Shop {

	//Configure the notifyManager function
	mgr := manager.GetManagerInstanceDefault()
	notifyFunc := mgr.HandleEventNotification

	//Select events to observe randomly
	var events []string
	for k := range HolidayEvents {
		events = append(events, k)
	}

	rand.Seed(time.Now().UnixNano())
	observedEvents := make([]string, numberofEvents)
	for i := range observedEvents {
		observedEvents[i] = events[rand.Intn(len(events))]
	}

	instance := &Shop{
		shopID:         shopID,
		location:       location,
		inventory:      make(map[string]*product.Product),
		numberofEvents: numberofEvents,
		observedEvents: observedEvents,
		notifyManager:  notifyFunc,
	}
	//Add the shop to the list of all shops
	allShops = append(allShops, instance)

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

// Add a method to append an event to the observedEvents
func (s *Shop) AddEvent(event string) {
	s.observedEvents = append(s.observedEvents, event)
	s.numberofEvents++
}

//TODO: Add a method to check all the stock of the products in the shop and trigger the replenishment function based on each product's replenishment rate
//Every day, the shop will loop through all the products in the shop and trigger the replenishment function

func (s *Shop) CheckAndNotify(date time.Time) {
	defer wg.Done()
	eventOccurred := false
	var eventToNotify string
	for _, event := range s.observedEvents {
		if eventDate, ok := HolidayEvents[event]; ok {
			if eventDate.Month() == date.Month() && eventDate.Day() == date.Day() {
				//if the matching event appears
				eventOccurred = true
				eventToNotify = event
				break
			}
		}
	}
	//Events check terminated

	var response *manager.Response
	if eventOccurred {
		response = s.notifyManager(eventToNotify, s.inventory)
	} else {
		response = s.notifyManager("", s.inventory)
	}

	if response != nil && response.Error == nil {
		for name, quantity := range response.Replenishments {
			s.inventory[name].SetNumber(s.inventory[name].GetNumber() + quantity)
		}
	}
}
