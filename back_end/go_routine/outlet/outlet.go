package outlet

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/event"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
	"github.com/gorilla/websocket"
)

type Outlet struct {
	outletID            string
	location            string
	inventory           map[string]*product.Product
	numberOfEvents      int
	events              map[string]*event.Event
	notifyCentralHub    func(string, string, string, string, *event.Event, map[string]*product.Product) *centralhub.Response
	clientPreferences   string
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

		//Initialize the items in the inventory to 0
		for _, prod := range product.GlobalProducts {
			inventory[prod.GetProductID()] = product.NewProduct(prod.GetProductID(), 0, prod.GetReplenishmentRate(), prod.GetMax_stock())
		}

		// Use the indices to add products to the outlet's inventory.
		for _, idx := range indices {
			productCopy := product.GlobalProducts[idx]
			//because the Product type holds no reference type, so we can just copy the product
			inventory[productCopy.GetProductID()].SetNumber(productCopy.GetNumber())
		}

		// Create a new outlet with the defined ID, location, and inventory.
		NewOutlet(fmt.Sprintf("%d", i+1), outletLocations[i], 6, event.HolidayMaps[i], &inventory)
	}
}

var (
	startChan = make(chan struct{})
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func startHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// wait for the message of frontend
	_, _, err = conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	//
	close(startChan)
}

func (o *Outlet) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a websocket connection
	o.wsupgrader.CheckOrigin = func(r *http.Request) bool {
		// Here you should implement a more robust check to ensure that the
		// request is coming from allowed origins.
		return true // Be cautious with this, it allows all origins!
	}

	conn, err := o.wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %v", err)
		return
	}
	o.client = conn
	// Need to hangup the connection in the main function "/centralhub"
}

var (
	currentDate time.Time
	dateMutex   sync.Mutex
	once        sync.Once
	ticker      *time.Ticker
	allOutlets  []*Outlet
)

func INIT() { // Single Agent
	centralhub.InitializeHub()
	product.InstanceProducts()
	InstanceOutlets()

	http.HandleFunc("/start", startHandler)
	// Wait for the message of frontend
	<-startChan
	once.Do(func() {
		// Initialize the currentDate to 31st December 2023
		currentDate = time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
		ticker = time.NewTicker(time.Second * 60)
		product.InstanceProducts() //Initialize the products

		immediate := time.After(0)
		go func() {
			http.ListenAndServe(":8001", nil)
		}()
		go func() {
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					performDailyOperations()
				case <-immediate:
					performDailyOperations()
					if immediate != nil {
						immediate = nil
					}
				}
			}

		}()
	})
}

func performDailyOperations() {
	_allOutlets := make([]*Outlet, len(allOutlets))
	copy(_allOutlets, allOutlets)
	localWg := new(sync.WaitGroup)
	localWg.Add(len(_allOutlets))

	// Current date add 1 day
	currentDate = currentDate.AddDate(0, 0, 1)

	for _, outlet := range _allOutlets {
		go func(outlet *Outlet, currentDate time.Time) {
			defer localWg.Done()
			outlet.CheckAndNotify(currentDate)
		}(outlet, currentDate)
	}
	localWg.Wait()
}

func GetCurrentDate() time.Time {
	dateMutex.Lock()
	defer dateMutex.Unlock()
	return currentDate
}

func NewOutlet(outletID string, location string, numberOfEvents int, holidayEvents map[string]*event.Event, inventory *map[string]*product.Product) *Outlet {
	centralHubInstance := centralhub.GetHubInstanceDefault()
	notifyFunc := centralHubInstance.HandleEventNotification

	events := make(map[string]*event.Event)
	for eventname, eventdetails := range holidayEvents {
		events[eventname] = &event.Event{
			EventDate:        eventdetails.EventDate,
			EventDescription: eventdetails.EventDescription,
		}
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
	// Set the client preferences
	switch location {
	case "Paris":
		o.clientPreferences = "Strong preference for Manchego Cheese and Olive Oil, moderate demand for Baguette, minimal interest in Black Tea."
	case "Lyon":
		o.clientPreferences = "Strong preference for Manchego Cheese and Olive Oil, moderate demand for Baguette, minimal interest in Black Tea."
	case "Marseille":
		o.clientPreferences = "High interest in Olive Oil and Black Tea, moderate preference for Baguette, low demand for Manchego Cheese."
	case "Nice":
		o.clientPreferences = "Strong demand for Olive Oil and Baguette, moderate interest in Black Tea, minimal preference for Manchego Cheese."
	}
	//copy the inventory
	for name, prod := range *inventory {
		o.inventory[name] = prod
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

func (o *Outlet) IntegrateResponseToSupermarketInfo(eventName string, Response *centralhub.Response) *centralhub.SupermarketInfo {
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
	supermarketInfo.Event = eventName
	return &supermarketInfo
}

func (o *Outlet) SendSupermarketInfoToFrontend(supermarketInfo *centralhub.SupermarketInfo) {

	//check the o.client is nil or not
	if o.client == nil {
		log.Println("Failed to send supermarket info to frontend: o.client is nil")
		return
	}
	// Send the json data to frontend via WebSocket
	err := o.client.WriteJSON(supermarketInfo)
	if err != nil {
		log.Printf("Failed to write json to frontend: %v", err)
	}
}

func (o *Outlet) AddProduct(p *product.Product) {
	// deepcopy
	newProduct := *p
	o.inventory[newProduct.GetProductID()] = &newProduct
}

func (o *Outlet) CheckAndNotify(date time.Time) {
	eventOccurred := false
	var eventName string
	var eventDetails *event.Event
	for eventName, eventDetails = range o.events {
		if eventDetails.EventDate.Month() == date.Month() && eventDetails.EventDate.Day() == date.Day() {
			//if the matching event appears
			eventOccurred = true
			break
		}
	}

	var response *centralhub.Response
	if eventOccurred {
		eventDetails.EventDate = date
		response = o.notifyCentralHub(o.GetOutletID(), o.GetLocation(), o.clientPreferences, eventName, eventDetails, o.inventory)
	} else {
		eventName = "No event"
		eventDetails.EventDate = date
		eventDetails.EventDescription = "No event"
		response = o.notifyCentralHub(o.GetOutletID(), o.GetLocation(), o.clientPreferences, eventName, eventDetails, o.inventory)
	}

	//Send the json pack SupermarketInfo to frontend
	o.SendSupermarketInfoToFrontend(o.IntegrateResponseToSupermarketInfo(eventName, response))
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

//End of Various methods to deal with delayed delivery
