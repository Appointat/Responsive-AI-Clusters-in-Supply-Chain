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
	// 创建Shop实例
	shop1 := shop.NewShop("Shop1", "Location1", 3)
	shop1.AddEvent("Test2")
	// 打印Shop实例以确认创建
	fmt.Println("Created shops:", shop1.GetShopID())

	time.Sleep(10 * time.Second)

	// 在测试结束后关闭计时器，可选
	// shop.StopTicker() // 假设您在shop包中有这样一个方法
}
