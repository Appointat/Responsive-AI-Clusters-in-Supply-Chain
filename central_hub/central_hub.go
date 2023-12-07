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
	Quantity          int    `json:"quantity"`
	ReplenishmentRate int    `json:"replenishmentRate"`
	MaxStock          int    `json:"maxStock"`
}

type AIRequest struct {
	outletID           string                 `json:"outletID"`
	outletlocation     string                 `json:"location"`
	centralhublocation string                 `json:"centralHubLocation"`
	date               time.Time              `json:"date"`
	Event              string                 `json:"event"`
	EventDescription   string                 `json:"eventDescription"`
	clientPreferences  string                 `json:"clientPreferences"`
	ShopInventory      map[string]ProductInfo `json:"shopInventory"` //商店库存
}

/*
clientPreferences: str
	The client's preferences for the products. This is a string of the form
	"product1:quantity1,product2:quantity2,...". The client's preferences
	are used to determine which products to replenish first.

centralHubInventory: dict
	A dictionary mapping product names to the number of products in the
	central hub's inventory.
*/

/*
***************************************************
To be combined with type Response in central_hub.go
***************************************************
*/
type AIResponse struct {
	DelayDays         int            `json:"delayDays"`
	CentralhubStock   map[string]int `json:"centralhubStock"`
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
	go http.ListenAndServe(":8080", nil)
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
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	h.client = conn
	// defer conn.Close()
	// Need to hangup the connection in the main function "/outlet+number"
}

// HandleEventNotification responds to an event by determining necessary product replenishments
func (h *CentralHub) HandleEventNotification(event string, date time.Time, shopInventory map[string]*product.Product) *Response {
	h.accessLock.Lock()
	defer h.accessLock.Unlock()
	inventoryInfo := make(map[string]ProductInfo)
	for _, prod := range shopInventory {
		inventoryInfo[prod.GetProductID()] = ProductInfo{
			ProductID:         prod.GetProductID(),
			Quantity:          prod.GetNumber(),
			ReplenishmentRate: prod.GetReplenishmentRate(),
			MaxStock:          prod.GetMax_stock(),
		}
	}

	// //Create Request
	// requestData := AIRequest{
	// 	Event:         event,
	// 	ShopInventory: inventoryInfo,
	// }

	// //Send Request to AI
	// aiResponse, err := h.SendRequestToAI(requestData)
	// if err != nil {
	// 	return &Response{
	// 		Replenishments: nil,
	// 		DeliveryTime:   0,
	// 		Error:          err,
	// 	}
	// }
	//暂时注释AI部分
	//Got Response from AI, send the general information to the frontend

	/*
		Test part
		We need to declare a template of GeneralInfo without AIResponse
	*/
	//Randomly generate the number of products in the warehouse
	WarehouseProduct := make(map[string]int)
	WarehouseProduct["Olive Oil"] = 1000
	WarehouseProduct["Baguette"] = 2000
	WarehouseProduct["Manchego Cheese"] = 1500
	WarehouseProduct["Black Tea"] = 800
	//Create a template of GeneralInfo
	generalInfo := GeneralInfo{
		Date:             date,
		Event:            event,
		WarehouseProduct: WarehouseProduct,
	}

	h.sendGeneralInfoToFrontEnd(&generalInfo)
	//Extract the info from aiResponse
	replenishments := make(map[string]int)

	// Calculate the number of products that need to be replenished
	// for name, _ := range shopInventory {
	// 	if ReplenishmentData, exists := aiResponse.ReplenishmentData[name]; exists {
	// 		quantityNeeded := ReplenishmentData
	// 		if quantityNeeded != 0 {
	// 			replenishments[name] = quantityNeeded
	// 			//decrease the stock of the product in the central hub
	// 			h.resources[name].SetNumber(h.resources[name].GetNumber() + quantityNeeded)
	// 		} else {
	// 			replenishments[name] = 0
	// 		}
	// 	}
	// }

	/*********************************************************************************************************************
	To be deleted after integration with AI
	*********************************************************************************************************************/
	if event != "" {
		fmt.Printf("Central Hub %s at %s received notification of event %s\n", h.hubID, h.location, event)
	}

	return &Response{
		Replenishments: replenishments,
		DeliveryTime:   0, //aiResponse.DelayDays
		Error:          nil,
	}
}
