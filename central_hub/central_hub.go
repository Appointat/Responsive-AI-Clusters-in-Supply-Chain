package centralhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/product"
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

type AIResponse struct {
	Reply             string                 `json:"reply"`
	ReplenishmentData map[string]ProductInfo `json:"replenishmentDates"`
}

/*
***************************************************
To be combined with type Response in central_hub.go
***************************************************
*/
type Response struct {
	Replenishments map[string]int // Key is the name of the product, value is the number of products to replenish
	Error          error
}

type CentralHub struct {
	hubID      string
	location   string
	resources  map[string]*product.Product
	accessLock sync.Mutex
}

var (
	instance *CentralHub
	once     sync.Once
)

// Singleton Pattern
func GetHubInstanceDefault() *CentralHub {
	once.Do(func() {
		instance = &CentralHub{}
	})
	return instance
}

// Initialize with Values
func GetHubInstance(hubID string, location string) *CentralHub {
	once.Do(func() {
		instance = &CentralHub{
			hubID:     hubID,
			location:  location,
			resources: make(map[string]*product.Product),
		}
	})
	return instance
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
	resp, err := http.Post("http://localhost:5000/ai", "application/json", bytes.NewBuffer(jsonData))
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

// HandleEventNotification responds to an event by determining necessary product replenishments
func (h *CentralHub) HandleEventNotification(event string, shopInventory map[string]*product.Product) *Response {
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
			Error:          err,
		}
	}

	//
	replenishments := make(map[string]int)

	// Calculate the number of products that need to be replenished
	for name, _ := range shopInventory {
		if replenishInfo, exists := aiResponse.ReplenishmentData[name]; exists {
			quantityNeeded := replenishInfo.Quantity
			if quantityNeeded != 0 {
				replenishments[name] = quantityNeeded
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
		Error:          nil,
	}
}
