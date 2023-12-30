package centralhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/event"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
	"github.com/gorilla/websocket"
)

type ProductInfo struct {
	Quantity          int `json:"current_storage_amount"`
	ReplenishmentRate int `json:"daily_replenishment_without_envent_from_central_hub"`
	MaxStock          int `json:"max_warehouse_capacity"`
}

type AIRequest struct {
	OutletID           string                 `json:"outlet_id"`
	Outletlocation     string                 `json:"outlet_location"`
	Centralhublocation string                 `json:"central_hub_location"`
	Date               time.Time              `json:"date"`
	Event              string                 `json:"event"`
	EventDescription   string                 `json:"event_description"`
	ClientPreferences  string                 `json:"client_preferences"`
	ShopInventory      map[string]ProductInfo `json:"outlet_inventory"`
}

// AIResponse is the response from the AI
type InventoryItem struct {
	CurrentStorageAmount int `json:"current_storage_amount"`
}

type ReplenishmentItem struct {
	ChangedReplenishmentAmount int    `json:"future_storage_amount"`
	UnusedReason               string `json:"specific_reason_of_replenishment"`
}

type AIResponse struct {
	DelayDays         int                          `json:"transportation_duration"`
	CentralhubStock   map[string]InventoryItem     `json:"central_hub_inventory"`
	ReplenishmentData map[string]ReplenishmentItem `json:"outlet_inventory"`
}

// GeneralInfo, SupermarketInfo, interact with the Frontend
type GeneralInfo struct {
	Date             time.Time      `json:"date"`
	WarehouseProduct map[string]int `json:"warehouseProduct"`
}

type SupermarketInfo struct {
	ID           string         `json:"id"`
	ProductLeft  map[string]int `json:"productLeft"` // The number of products left in the supermarket
	ProductAdd   map[string]int `json:"productAdd"`  // The number of products added to the supermarket
	DeliveryTime int            `json:"deliveryTime"`
	Event        string         `json:"event"`
}

// End of the part to integrate with Frontend
type Response struct {
	Replenishments map[string]int // Key is the name of the product, value is the number of products to replenish
	DeliveryTime   int
	Error          error
}

type CentralHub struct {
	hubID       string
	location    string
	resources   map[string]*product.Product
	accessLock  sync.Mutex
	aiServerURL string
	wsupgrader  websocket.Upgrader
	client      *websocket.Conn
}

var (
	instance *CentralHub
	once     sync.Once
)

// Singleton Pattern
func GetHubInstanceDefault() *CentralHub {
	once.Do(func() {
		instance = &CentralHub{
			aiServerURL: "http://localhost:5000/ai",
		}
	})
	return instance
}

// Initialize with Values
func GetHubInstance(hubID string, location string) *CentralHub {
	once.Do(func() {
		instance = &CentralHub{
			hubID:       hubID,
			location:    location,
			resources:   make(map[string]*product.Product),
			aiServerURL: "http://localhost:5000/ai",
		}
	})
	return instance
}
func (h *CentralHub) SetInventory(inventory map[string]*product.Product) {
	h.resources = inventory
}
func InitializeHub() {
	GetHubInstance("Central Hub", "Paris")
	inventory := make(map[string]*product.Product)

	// Initialize the inventory
	inventory["Olive Oil"] = product.NewProduct("Olive Oil", 1000, 30, 500)
	inventory["Baguette"] = product.NewProduct("Baguette", 2000, 50, 300)
	inventory["Manchego Cheese"] = product.NewProduct("Manchego Cheese", 1500, 40, 400)
	inventory["Black Tea"] = product.NewProduct("Black Tea", 800, 20, 250)
	http.HandleFunc("/centralhub", instance.HandleWebSocket)
	go http.ListenAndServe(":8001", nil)
	// Add the inventory to the central hub
	instance.SetInventory(inventory)
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

func (h *CentralHub) SendRequestToAI(requestData AIRequest) (*AIResponse, error) {
	// Code as JSON

	jsonData, _ := json.Marshal(requestData)

	// Send Post Request to Python AI
	resp, err := http.Post(h.aiServerURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Error occurred in communication with AI: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	var aiResponse AIResponse
	json.Unmarshal([]byte(jsonData), &aiResponse)

	// If the AI response is nil, return an error
	if aiResponse.CentralhubStock == nil || aiResponse.ReplenishmentData == nil || aiResponse.DelayDays == 0 {
		log.Printf("AI Response is nil\n")
		log.Printf("AI Response: %v\n", aiResponse)
		return nil, fmt.Errorf("AI Response is nil")
	}

	log.Printf("AI Response: %v\n", aiResponse)

	return &aiResponse, nil
}

// Send the general information to the frontend
func (h *CentralHub) IntegrateAIResponseToGeneralInfo(event string, date time.Time, aiResponseData *AIResponse) *GeneralInfo {
	// Integrate the data from AI

	// Extract the info from aiResponse
	warehouseProduct := make(map[string]int)
	for name, item := range aiResponseData.CentralhubStock {
		warehouseProduct[name] = item.CurrentStorageAmount
	}
	generalInfo := GeneralInfo{
		Date: date,
		// Event:            event,
		WarehouseProduct: warehouseProduct,
	}

	return &generalInfo
}

func (h *CentralHub) sendGeneralInfoToFrontEnd(info *GeneralInfo) {
	// Send the general information to the frontend

	//check the h.client is nil or not
	if h.client == nil {
		fmt.Println("Error in sending general info to front end: h.client is nil")
		return
	}
	// Send Post Request to Frontend
	err := h.client.WriteJSON(info)
	if err != nil {
		log.Printf("Error sending message to WebSocket: %v\n", err)
	}
}

func (h *CentralHub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	h.wsupgrader.CheckOrigin = func(r *http.Request) bool {
		// Here you should implement a more robust check to ensure that the
		// request is coming from allowed origins.
		return true // Be cautious with this, it allows all origins!
	}

	// Upgrade the HTTP connection to a websocket connection
	conn, err := h.wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %v", err)
		return
	}
	h.client = conn
	// defer conn.Close()
	// Need to hangup the connection in the main function "/outlet+number"
}

// HandleEventNotification responds to an event by determining necessary product replenishments
func (h *CentralHub) HandleEventNotification(outletID string, outletlocation string, clientPreferences string, eventName string, event *event.Event, shopInventory map[string]*product.Product) *Response {
	h.accessLock.Lock()
	defer h.accessLock.Unlock()
	inventoryInfo := make(map[string]ProductInfo)
	for _, prod := range shopInventory {
		inventoryInfo[prod.GetProductID()] = ProductInfo{
			Quantity:          prod.GetNumber(),
			ReplenishmentRate: prod.GetReplenishmentRate(),
			MaxStock:          prod.GetMax_stock(),
		}
	}

	// //Create Request
	requestData := AIRequest{
		OutletID:           outletID,
		Outletlocation:     outletlocation,
		Centralhublocation: h.location,
		Date:               event.EventDate,
		Event:              eventName,
		EventDescription:   event.EventDescription,
		ClientPreferences:  clientPreferences,
		ShopInventory:      inventoryInfo,
	}

	// //Send Request to AI
	var aiResponse *AIResponse
	aiResponse, err := h.SendRequestToAI(requestData)
	if err != nil {
		log.Printf("Error occurred in communication with AI: %v\n", err)
		return &Response{
			Replenishments: nil,
			DeliveryTime:   0,
			Error:          err,
		}
	}

	h.sendGeneralInfoToFrontEnd(h.IntegrateAIResponseToGeneralInfo(eventName, event.EventDate, aiResponse))
	//Extract the info from aiResponse
	replenishments := make(map[string]int) // TODO: Extract the replenishment info from aiResponse
	//Calculate the number of products that need to be replenished
	for name, _ := range shopInventory {
		//Switch name keys to _
		switch name {
		case "Baguette":
			name = "baguette"
		case "Black Tea":
			name = "black_tea"
		case "Manchego Cheese":
			name = "manchego_cheese"
		case "Olive Oil":
			name = "olive_oil"
		}

		if ReplenishmentData, exists := aiResponse.ReplenishmentData[name]; exists {
			//change back
			switch name {
			case "baguette":
				name = "Baguette"
			case "black_tea":
				name = "Black Tea"
			case "manchego_cheese":
				name = "Manchego Cheese"
			case "olive_oil":
				name = "Olive Oil"
			}

			quantityNeeded := ReplenishmentData.ChangedReplenishmentAmount
			if quantityNeeded != 0 {
				replenishments[name] = quantityNeeded
				//Directly change the number of products in the central hub according to the AI response
				h.resources[name].SetNumber(aiResponse.CentralhubStock[name].CurrentStorageAmount)
			} else {
				replenishments[name] = 0
			}
		}
	}

	/*********************************************************************************************************************
	To be deleted after integration with AI
	*********************************************************************************************************************/
	if event != nil {
		fmt.Printf("Central Hub %s at %s received notification of event %s\n", h.hubID, h.location, event)
	}

	return &Response{
		Replenishments: replenishments,
		DeliveryTime:   aiResponse.DelayDays,
		Error:          nil,
	}
}
