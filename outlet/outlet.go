package outlet

import (
	"math/rand"
	"sync"
	"time"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

// HolidayEvents maps event names to their dates.
var HolidayEvents = map[string]time.Time{
	"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
	"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
	// ... (other events) ...
	"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
	"Test2":                time.Date(2023, 11, 30, 0, 0, 0, 0, time.UTC),
}

type Outlet struct {
	outletID         string
	location         string
	inventory        map[string]*product.Product
	numberOfEvents   int
	observedEvents   []string
	notifyCentralHub func(string, map[string]*product.Product) *centralhub.Response
}

var (
	currentDate time.Time
	dateMutex   sync.Mutex
	once        sync.Once
	wg          sync.WaitGroup
	ticker      *time.Ticker
	allOutlets  []*Outlet
)

func init() {
	// Initialize the HolidayEvents with the current year.
	currentYear := time.Now().Year()
	for k, v := range HolidayEvents {
		HolidayEvents[k] = time.Date(currentYear, v.Month(), v.Day(), 0, 0, 0, 0, time.UTC)
	}

	once.Do(func() {
		currentDate = time.Now()
		ticker = time.NewTicker(time.Second * 1)
		go func() {
			for range ticker.C {
				dateMutex.Lock()
				currentDate = currentDate.AddDate(0, 0, 1)
				dateMutex.Unlock()

				wg.Add(len(allOutlets))
				for _, o := range allOutlets {
					go o.CheckAndNotify(currentDate)
				}

				wg.Wait()
			}
		}()
	})
}

func GetCurrentDate() time.Time {
	dateMutex.Lock()
	defer dateMutex.Unlock()
	return currentDate
}

func NewOutlet(outletID, location string, numberOfEvents int) *Outlet {
	centralHubInstance := centralhub.GetHubInstanceDefault()
	notifyFunc := centralHubInstance.HandleEventNotification

	events := make([]string, 0, len(HolidayEvents))
	for k := range HolidayEvents {
		events = append(events, k)
	}

	rand.Seed(time.Now().UnixNano())
	observedEvents := make([]string, numberOfEvents)
	for i := 0; i < numberOfEvents; i++ {
		observedEvents[i] = events[rand.Intn(len(events))]
	}

	o := &Outlet{
		outletID:         outletID,
		location:         location,
		inventory:        make(map[string]*product.Product),
		numberOfEvents:   numberOfEvents,
		observedEvents:   observedEvents,
		notifyCentralHub: notifyFunc,
	}

	allOutlets = append(allOutlets, o)
	return o
}

func (o *Outlet) GetOutletID() string {
	return o.outletID
}

func (o *Outlet) GetLocation() string {
	return o.location
}

func (o *Outlet) GetNumberOfProducts(productName string) int {
	if p, ok := o.inventory[productName]; ok {
		return p.GetNumber()
	}
	return -1
}

func (o *Outlet) AddEvent(event string) {
	o.observedEvents = append(o.observedEvents, event)
}

func (o *Outlet) CheckAndNotify(date time.Time) {
	defer wg.Done()

	eventOccurred := false
	var eventToNotify string

	for _, event := range o.observedEvents {
		if eventDate, ok := HolidayEvents[event]; ok {
			if eventDate.Month() == date.Month() && eventDate.Day() == date.Day() {
				//if the matching event appears
				eventOccurred = true
				eventToNotify = event
				break
			}
		}
	}

	var response *centralhub.Response
	if eventOccurred {
		response = o.notifyCentralHub(eventToNotify, o.inventory)
	} else {
		response = o.notifyCentralHub("", o.inventory)
	}

	if response != nil && response.Error == nil {
		for name, quantity := range response.Replenishments {
			if prod, ok := o.inventory[name]; ok {
				prod.SetNumber(prod.GetNumber() + quantity)
			}
		}
	}
}
