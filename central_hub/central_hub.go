package centralhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
	"github.com/gorilla/websocket"
)

type ProductInfo struct {
	ProductID         string `json:"productID"`
	ProductName       string `json:"productName"`
	Quantity          int    `json:"quantity"`
	ReplenishmentRate int    `json:"replenishmentRate"`
	MaxStock          int    `json:"maxStock"`
}

type AIRequest struct {
	Event         string                 `json:"event"`
	ShopInventory map[string]ProductInfo `json:"shopInventory"`
}

/*
***************************************************
To be combined with type Response in central_hub.go
***************************************************
*/
type AIResponse struct {
	Reply             string         `json:"reply"`
	DelayDays         int            `json:"delayDays"`
	ReplenishmentData map[string]int `json:"replenishmentDates"`
}

// GeneralInfo, SupermarketInfo, interact with the Frontend
type GeneralInfo struct {
	Date             time.Time      `json:"date"`
	Event            string         `json:"event"`
	WarehouseProduct map[string]int `json:"warehouseProduct"`
}

type SupermarketInfo struct {
	ID           string         `json:"id"`
	ProductLeft  map[string]int `json:"productLeft"` // The number of products left in the supermarket
	ProductAdd   map[string]int `json:"productAdd"`  // The number of products added to the supermarket
	DeliveryTime int            `json:"deliveryTime"`
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
	// 将传入的库存映射赋值给中心仓库实例的Inventory属性
	h.resources = inventory
}
func InitializeHub() {
	GetHubInstance("Central Hub", "Paris")
	inventory := make(map[string]*product.Product)

	// Initialize the inventory
	inventory["P001"] = product.NewProduct("A", "Olive Oil", 1000, 30, 500)
	inventory["P002"] = product.NewProduct("B", "Baguette", 2000, 50, 300)
	inventory["P003"] = product.NewProduct("C", "Manchego Cheese", 1500, 40, 400)
	inventory["P004"] = product.NewProduct("D", "Black Tea", 800, 20, 250)

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
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	// Send Post Request to Python AI
	resp, err := http.Post(h.aiServerURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode JSON Response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convert JSON to Struct
	var response AIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Send the general information to the frontend
func (h *CentralHub) IntegrateAIResponseToGeneralInfo(event string, date time.Time, aiResponseData *AIResponse) *GeneralInfo {
	// Integrate the data from AI

	// Extract the info from aiResponse
	warehouseProduct := aiResponseData.ReplenishmentData
	generalInfo := GeneralInfo{
		Date:             date,
		Event:            event,
		WarehouseProduct: warehouseProduct,
	}

	return &generalInfo
}

func (h *CentralHub) sendGeneralInfoToFrontEnd(info *GeneralInfo) {
	// Send the general information to the frontend

	// Code as JSON
	jsonData, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error in sending general info to front end:", err)
		return
	}

	// Send Post Request to Frontend
	err = h.client.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Printf("Error sending message to WebSocket: %v\n", err)
	}
}

func (h *CentralHub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a websocket connection
	conn, err := h.wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	h.client = conn
	defer conn.Close()
	// Need to hangup the connection in the main function "/outlet+number"
}

// HandleEventNotification responds to an event by determining necessary product replenishments
func (h *CentralHub) HandleEventNotification(event string, date time.Time, shopInventory map[string]*product.Product) *Response {
	h.accessLock.Lock()
	defer h.accessLock.Unlock()
	inventoryInfo := make(map[string]ProductInfo)
	for _, prod := range shopInventory {
		inventoryInfo[prod.GetProductName()] = ProductInfo{
			ProductID:         prod.GetProductID(),
			ProductName:       prod.GetProductName(),
			Quantity:          prod.GetNumber(),
			ReplenishmentRate: prod.GetReplenishmentRate(),
			MaxStock:          prod.GetMax_stock(),
		}
	}

	//Create Request
	requestData := AIRequest{
		Event:         event,
		ShopInventory: inventoryInfo,
	}

	//Send Request to AI
	aiResponse, err := h.SendRequestToAI(requestData)
	if err != nil {
		return &Response{
			Replenishments: nil,
			DeliveryTime:   0,
			Error:          err,
		}
	}
	//Got Response from AI, send the general information to the frontend
	h.sendGeneralInfoToFrontEnd(h.IntegrateAIResponseToGeneralInfo(event, date, aiResponse))
	//Extract the info from aiResponse
	replenishments := make(map[string]int)

	// Calculate the number of products that need to be replenished
	for name, _ := range shopInventory {
		if ReplenishmentData, exists := aiResponse.ReplenishmentData[name]; exists {
			quantityNeeded := ReplenishmentData
			if quantityNeeded != 0 {
				replenishments[name] = quantityNeeded
				//decrease the stock of the product in the central hub
				h.resources[name].SetNumber(h.resources[name].GetNumber() + quantityNeeded)
			} else {
				replenishments[name] = 0
			}
		}
	}

	/*********************************************************************************************************************
	To be deleted after integration with AI
	*********************************************************************************************************************/
	if event != "" {
		fmt.Printf("Central Hub %s at %s received notification of event %s\n", h.hubID, h.location, event)
	}

	return &Response{
		Replenishments: replenishments,
		DeliveryTime:   aiResponse.DelayDays,
		Error:          nil,
	}
}
