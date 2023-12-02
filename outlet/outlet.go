package outlet

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
	"github.com/gorilla/websocket"
)

type Outlet struct {
	outletID            string
	location            string
	inventory           map[string]*product.Product
	numberOfEvents      int
	events              map[string]time.Time
	notifyCentralHub    func(string, time.Time, map[string]*product.Product) *centralhub.Response
	scheduledDeliveries map[time.Time]map[string]int
	//websocket connection
	wsupgrader websocket.Upgrader
	client     *websocket.Conn
}

func InstanceOutlets() {
	// Define the unique combination of product indices for each outlet.
	outletProductIndices := [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{1, 2, 3},
	}

	// Define the locations for each outlet.
	outletLocations := []string{"Paris", "Lyon", "Marseille", "Nice"}

	//clear the allOutlets
	allOutlets = []*Outlet{}
	// Loop through the defined configurations and create outlets.
	for i, indices := range outletProductIndices {
		inventory := make(map[string]*product.Product)

		// Use the indices to add products to the outlet's inventory.
		for _, idx := range indices {
			product := product.GlobalProducts[idx]
			productCopy := product
			//because the Product type holds no reference type, so we can just copy the product
			inventory[productCopy.GetProductID()] = &productCopy
		}

		// Create a new outlet with the defined ID, location, and inventory.
		allOutlets = append(allOutlets, &Outlet{
			outletID:  fmt.Sprintf("%d", i+1), // Outlet IDs are "1", "2", "3", "4".
			location:  outletLocations[i],
			inventory: inventory,
		})
	}
}

func (o *Outlet) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a websocket connection
	conn, err := o.wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	o.client = conn
	defer conn.Close()
	// Need to hangup the connection in the main function "/centralhub"
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
		product.InstanceProducts() //Initialize the products

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
					go func(outlet *Outlet, currentDate time.Time) {
						defer localWg.Done()
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

func NewOutlet(outletID string, location string, numberOfEvents int, holidayEvents map[string]time.Time) *Outlet {
	centralHubInstance := centralhub.GetHubInstanceDefault()
	notifyFunc := centralHubInstance.HandleEventNotification

	events := make(map[string]time.Time)
	for event, eventDate := range holidayEvents {
		events[event] = eventDate
	}

	o := &Outlet{
		outletID:            outletID,
		location:            location,
		inventory:           make(map[string]*product.Product), // TODO: initialize with products
		numberOfEvents:      numberOfEvents,
		events:              events,
		notifyCentralHub:    notifyFunc,
		scheduledDeliveries: make(map[time.Time]map[string]int),
	}

	allOutlets = append(allOutlets, o)

	//attach the websocket connection to the outlet
	route := fmt.Sprintf("/outlet%d", len(allOutlets))
	http.HandleFunc(route, o.HandleWebSocket)
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

func (o *Outlet) IntegrateResponseToSupermarketInfo(Response *centralhub.Response) *centralhub.SupermarketInfo {
	//Extract the replenishment information from the response into the supermarketInfo
	var supermarketInfo centralhub.SupermarketInfo
	supermarketInfo.ProductAdd = make(map[string]int)
	for name, quantity := range Response.Replenishments {
		supermarketInfo.ProductAdd[name] = quantity
	}
	supermarketInfo.DeliveryTime = Response.DeliveryTime
	supermarketInfo.ID = o.outletID

	//Extract the inventory information from the outlet into the supermarketInfo
	supermarketInfo.ProductLeft = make(map[string]int)
	for name, product := range o.inventory {
		supermarketInfo.ProductLeft[name] = product.GetNumber()
	}
	return &supermarketInfo
}

func (o *Outlet) SendSupermarketInfoToFrontend(supermarketInfo *centralhub.SupermarketInfo) {
	//Send the json pack SupermarketInfo to frontend
	err := o.client.WriteJSON(supermarketInfo)
	if err != nil {
		log.Println("Failed to write json to frontend: %+v", err)
	}
}

func (o *Outlet) AddProduct(p *product.Product) {
	// deepcopy
	newProduct := *p
	o.inventory[newProduct.GetProductID()] = &newProduct
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
		response = o.notifyCentralHub(eventToNotify, date, o.inventory)
	} else {
		response = o.notifyCentralHub("", date, o.inventory)
	}
	//Send the json pack SupermarketInfo to frontend
	o.SendSupermarketInfoToFrontend(o.IntegrateResponseToSupermarketInfo(response))
	//Process the scheduled deliveries
	o.ProcessScheduledDeliveries(date)
	o.scheduleDeliveries(response, date)

}

// Various methods to deal with delayed delivery
func (o *Outlet) scheduleDeliveries(response *centralhub.Response, currentDate time.Time) {
	for name, quantity := range response.Replenishments {
		deliveryDate := currentDate.Add(time.Duration(response.DeliveryTime) * time.Hour * 24)

		// if today is the delivery date or the delivery time is 0, then deliver immediately
		if deliveryDate.Equal(currentDate) || response.DeliveryTime == 0 {
			if prod, ok := o.inventory[name]; ok {
				prod.SetNumber(prod.GetNumber() + quantity)
			}
		} else {
			// else schedule the delivery
			if _, exists := o.scheduledDeliveries[deliveryDate]; !exists {
				o.scheduledDeliveries[deliveryDate] = make(map[string]int)
			}
			o.scheduledDeliveries[deliveryDate][name] = quantity
		}
	}
}

func (o *Outlet) ProcessScheduledDeliveries(currentDate time.Time) {
	if deliveries, exists := o.scheduledDeliveries[currentDate]; exists {
		for name, quantity := range deliveries {
			if p, ok := o.inventory[name]; ok {
				p.SetNumber(p.GetNumber() + quantity)
			}
		}
		delete(o.scheduledDeliveries, currentDate)
	}
}

func (o *Outlet) updateInventory(replenishments map[string]int) {
	for name, quantity := range replenishments {
		if prod, ok := o.inventory[name]; ok {
			prod.SetNumber(prod.GetNumber() + quantity)
		}
	}
}

//End of Various methods to deal with delayed delivery
