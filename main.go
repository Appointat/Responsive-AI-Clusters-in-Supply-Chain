package main

import (
	"fmt"
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/manager"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/shop"
)

func main() {
	//create a manager
	manager.GetManagerInstance("Manager1", "Location1")
	shop1 := shop.NewShop("Shop1", "Location1", 3)
	shop1.AddEvent("Test2")
	fmt.Println("Created shops:", shop1.GetShopID())
	time.Sleep(10 * time.Second)
}
