package outlet

import (
	"sync"
	"time"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
)

type Outlet struct {
	outletID         string
	location         string
	inventory        map[string]*product.Product
	numberOfEvents   int
	events           map[string]time.Time
	notifyCentralHub func(string, map[string]*product.Product) *centralhub.Response
}

var (
	initDate    time.Time
	currentDate time.Time
	dateMutex   sync.Mutex
	once        sync.Once
	ticker      *time.Ticker
	allOutlets  []*Outlet
)

func init() { // Single Agent
	once.Do(func() {
		initDate = time.Now()
		ticker = time.NewTicker(time.Second * 1)

		go func() {
			defer ticker.Stop()
			for range ticker.C {
				_allOutlets := make([]*Outlet, len(allOutlets))
				copy(_allOutlets, allOutlets)

				localWg := new(sync.WaitGroup)
				localWg.Add(len(_allOutlets))

				for _, outlet := range _allOutlets {
					// Get the virtual current date
					diff := time.Since(initDate)
					seconds := int(diff.Seconds())
					factor := 3600 * 24 // 1 day
					virtualSeconds := seconds * factor
					currentDate := initDate.Add(time.Second * time.Duration(virtualSeconds))
					//send CURRENTDATE to front
					go func(outlet *Outlet, currentDate time.Time) {
						defer localWg.Done()
						/************************
						SEND CURRENTDATE TO FRONT
						************************/
						outlet.CheckAndNotify(currentDate)
					}(outlet, currentDate)
				}
				localWg.Wait()
			}
		}()
	})
}

func GetCurrentDate() time.Time {
	dateMutex.Lock()
	defer dateMutex.Unlock()
	return currentDate
}

func NewOutlet(outletID, location string, numberOfEvents int, holidayEvents map[string]time.Time) *Outlet {
	centralHubInstance := centralhub.GetHubInstanceDefault()
	notifyFunc := centralHubInstance.HandleEventNotification

	events := make(map[string]time.Time)
	for event, eventDate := range holidayEvents {
		events[event] = eventDate
	}

	o := &Outlet{
		outletID:         outletID,
		location:         location,
		inventory:        make(map[string]*product.Product), // TODO: initialize with products
		numberOfEvents:   numberOfEvents,
		events:           events,
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

func (o *Outlet) CheckAndNotify(date time.Time) {
	eventOccurred := false
	var eventToNotify string

	for event, eventDate := range o.events {
		if eventDate.Year() == date.Year() && eventDate.Month() == date.Month() && eventDate.Day() == date.Day() {
			//if the matching event appears
			eventOccurred = true
			eventToNotify = event
			break
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

	// TODO: http.Post
}
