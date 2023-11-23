package main

import (
	"fmt"
	"log"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub" // 替换为您的包名
)

func main() {
	// 创建 CentralHub 实例
	hub := centralhub.GetHubInstance("Central_Hub1", "Location1")

	// 构造 AIRequest
	requestData := centralhub.AIRequest{
		Event: "test event",
		ShopInventory: map[string]centralhub.ProductInfo{
			"Product1": centralhub.ProductInfo{
				ProductID:         "P123",
				ProductName:       "Widget",
				Quantity:          10,
				ReplenishmentRate: 5,
				MaxStock:          100,
			},
		},
	}

	// 调用 SendRequestToAI 并打印结果
	response, err := hub.SendRequestToAI(requestData)
	if err != nil {
		log.Fatalf("SendRequestToAI returned an error: %v", err)
	}

	fmt.Printf("Response from AI: %+v\n", response)
}
